// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <math.h>
#include "wav.h"

int wav_sample_amp(int16_t *sample,int len,float factor);
int wav_getdb(double factor);

wav_t *wav_open(char *file_name){
    wav_t *wav = NULL;
    char buffer[1024];
    int read_len = 0;
    int offset = 0;


    if(NULL == file_name){
        printf("file_name is NULL\n");
        return NULL;
    }
    wav = (wav_t *)malloc(sizeof(wav_t));
    if(NULL == wav){
        printf("malloc wav failedly\n");
        return NULL;
    }
    memset(wav, 0,sizeof(wav_t));

    wav->fp = fopen(file_name, "rb+");
    if(NULL == wav->fp){
        printf("fopen %s failedly\n", file_name);
        free(wav);
        return NULL;
    }

    //handle RIFF WAVE chunk
    read_len = fread(buffer, 1, 12, wav->fp);
    if(read_len < 12){
        printf("error wav file\n");
        wav_close(&wav);
        return NULL;
    }
    if(strncasecmp("RIFF", buffer, 4)){
        printf("error wav file\n");
        wav_close(&wav);
        return NULL;
    }
    memcpy(wav->riff.id, buffer, 4);
    wav->riff.size = *(int *)(buffer + 4);
    if(strncasecmp("WAVE", buffer + 8, 4)){
        printf("error wav file\n");
        wav_close(&wav);
        return NULL;
    }
    memcpy(wav->riff.type, buffer + 8, 4);
    wav->file_size = wav->riff.size + 8;
    offset += 12;
    while(1){
        char id_buffer[5] = {0};
        int tmp_size = 0;

        read_len = fread(buffer, 1, 8, wav->fp);
        if(read_len < 8){
            printf("error wav file\n");
            wav_close(&wav);
            return NULL;
        }
        memcpy(id_buffer, buffer, 4);
        tmp_size = *(int *)(buffer + 4);


        if(0 == strncasecmp("FMT", id_buffer, 3)){
            memcpy(wav->format.id, id_buffer, 3);
            wav->format.size = tmp_size;
            read_len = fread(buffer, 1, tmp_size, wav->fp);
            if(read_len < tmp_size){
                printf("error wav file\n");
                wav_close(&wav);
                return NULL;
            }
            wav->format.compression_code = *(short *)buffer;
            wav->format.channels = *(short *)(buffer + 2);
            wav->format.samples_per_sec = *(int *)(buffer + 4);
            wav->format.avg_bytes_per_sec = *(int *)(buffer + 8);
            wav->format.block_align = *(short *)(buffer + 12);
            wav->format.bits_per_sample = *(short *)(buffer + 14);
        }
        else if(0 == strncasecmp("DATA", id_buffer, 4)){
            memcpy(wav->data.id, id_buffer, 4);
            wav->data.size = tmp_size;
            offset += 8;
            wav->data_offset = offset;
            wav->data_size = wav->data.size;
            break;
        }
        else{
            printf("unhandled chunk: %s, size: %d\n", id_buffer, tmp_size);
            fseek(wav->fp, tmp_size, SEEK_CUR);
        }
        offset += 8 + tmp_size;
    }
    return wav;
}

void wav_close(wav_t **wav){
    wav_t *tmp_wav;
    if(NULL == wav){
        return ;
    }

    tmp_wav = *wav;
    if(NULL == tmp_wav){
        return ;
    }


    if(NULL != tmp_wav->fp){
        fclose(tmp_wav->fp);
    }
    free(tmp_wav);

    *wav = NULL;
}


void wav_rewind(wav_t *wav){
    if(fseek(wav->fp, wav->data_offset, SEEK_SET) < 0){
        printf("wav rewind failedly\n");
    }
}


int wav_over(wav_t *wav){
    return feof(wav->fp);
}


int wav_read_data(wav_t *wav, char *buffer, int buffer_size){
    return fread(buffer, 1, buffer_size, wav->fp);
}


void wav_dump(wav_t *wav){
    printf("file length: %d\n", wav->file_size);


    printf("\nRIFF WAVE Chunk\n");
    printf("id: %s\n", wav->riff.id);
    printf("size: %d\n", wav->riff.size);
    printf("type: %s\n", wav->riff.type);


    printf("\nFORMAT Chunk\n");
    printf("id: %s\n", wav->format.id);
    printf("size: %d\n", wav->format.size);
    if(wav->format.compression_code == 0){
        printf("compression: Unknown\n");
    }
    else if(wav->format.compression_code == 1){
        printf("compression: PCM/uncompressed\n");
    }
    else if(wav->format.compression_code == 2){
        printf("compression: Microsoft ADPCM\n");
    }
    else if(wav->format.compression_code == 6){
        printf("compression: ITU G.711 a-law\n");
    }
    else if(wav->format.compression_code == 7){
        printf("compression: ITU G.711 ?μ-law\n");
    }
    else if(wav->format.compression_code == 17){
        printf("compression: IMA ADPCM\n");
    }
    else if(wav->format.compression_code == 20){
        printf("compression: ITU G.723 ADPCM (Yamaha)\n");
    }
    else if(wav->format.compression_code == 49){
        printf("compression: GSM 6.10\n");
    }
    else if(wav->format.compression_code == 64){
        printf("compression: ITU G.721 ADPCM\n");
    }
    else if(wav->format.compression_code == 80){
        printf("compression: MPEG\n");
    }
    else{
        printf("compression: Unknown\n");
    }

    printf("channels: %d\n", wav->format.channels);
    printf("samples: %d\n", wav->format.samples_per_sec);
    printf("avg_bytes_per_sec: %d\n", wav->format.avg_bytes_per_sec);
    printf("block_align: %d\n", wav->format.block_align);
    printf("bits_per_sample: %d\n", wav->format.bits_per_sample);


    printf("\nDATA Chunk\n");
    printf("id: %s\n", wav->data.id);
    printf("size: %d\n", wav->data.size);
    printf("data offset: %d\n", wav->data_offset);
}

int wav_write_header(wav_t *wav,FILE *fp){
    int write_len = 0;
    char tempbuff[36];
    char *pointer = tempbuff;
    strncpy(pointer,wav->riff.id,4);
    pointer += 4;
    strncpy(pointer,(char *)(&(wav->riff.size)),4);
    printf("riff size:%d.\n",wav->riff.size);
    pointer += 4;
    strncpy(pointer,wav->riff.type,4);
    pointer += 4;

    strncpy(pointer,wav->format.id,4);
    pointer += 4;
    strncpy(pointer,(char *)(&(wav->format.size)),4);
    pointer += 4;
    strncpy(pointer,(char *)(&(wav->format.compression_code)),2);
    pointer += 2;
    strncpy(pointer,(char *)(&(wav->format.channels)),2);
    pointer += 2;
    strncpy(pointer,(char *)(&(wav->format.samples_per_sec)),4);
    pointer += 4;
    strncpy(pointer,(char *)(&(wav->format.avg_bytes_per_sec)),4);
   pointer += 4;
    strncpy(pointer,(char *)(&(wav->format.block_align)),2);
    pointer += 2;
    strncpy(pointer,(char *)(&(wav->format.bits_per_sample)),2);
    pointer += 2;

    printf("hex:");
   for(int i=0;i<36;i++){
        printf("%02X ",tempbuff[i]);
    }
    printf("\n");

    write_len = fwrite(tempbuff, 1, 36, fp);
    printf("writelen:%d\n",write_len);
    return 0;
}

int wav_convert(char *file_name,char *dest_name,float precent){
    wav_t *wav = NULL;
    char buffer[1024];
    int read_len = 0;
    int write_len = 0;
    int offset = 0;
    FILE *dest_fp = NULL;

    if(NULL == file_name || dest_name == NULL){
        printf("file is NULL\n");
        return -1;
    }

    double factor = 0.0;
    if (precent > 0.00){
        //int db = -4;
        int db = wav_getdb(precent);
        printf("db: %d.\n",db);

        //注意：pow里必须强转为double类型，因为两个int类型做除法，结果还是int类型，会损失精度，此处把db强转为
        //double，double/int最后的结果是个double，保证了精度
        factor= pow(10, (double)db / 20);
    }
   
    wav = (wav_t *)malloc(sizeof(wav_t));
    if(NULL == wav){
        printf("malloc wav failedly\n");
        return -1;
    }
    memset(wav, 0,sizeof(wav_t));

    wav->fp = fopen(file_name, "rb");
    if(NULL == wav->fp){
        printf("fopen %s failedly\n", file_name);
        free(wav);
        return -1;
    }

   //handle RIFF WAVE chunk
    read_len = fread(buffer, 1, 12, wav->fp);
    if(read_len < 12){
        printf("error wav file\n");
        wav_close(&wav);
        return -1;
    }
    if(strncasecmp("RIFF", buffer, 4)){
        printf("error wav file\n");
        wav_close(&wav);
        return -1;
    }
    memcpy(wav->riff.id, buffer, 4);
    wav->riff.size = *(int *)(buffer + 4);
    if(strncasecmp("WAVE", buffer + 8, 4)){
        printf("error wav file\n");
        wav_close(&wav);
        return -1;
    }
    memcpy(wav->riff.type, buffer + 8, 4);
    wav->file_size = wav->riff.size + 8;
    offset += 12;

    dest_fp = fopen(dest_name, "wb+");
    if(NULL == dest_fp){
        printf("fopen %s failedly\n", file_name);
        wav_close(&wav);
        return -1;
    }

    while(1){
        char id_buffer[5] = {0};
        int tmp_size = 0;

        read_len = fread(buffer, 1, 8, wav->fp);
        if(read_len < 8){
            printf("error wav file\n");
            wav_close(&wav);
            fclose(dest_fp);
            return -1;
        }
        memcpy(id_buffer, buffer, 4);
        tmp_size = *(int *)(buffer + 4);

        if(0 == strncasecmp("FMT", id_buffer, 3)){
            memcpy(wav->format.id, id_buffer, 4);
            wav->format.size = tmp_size;
            read_len = fread(buffer, 1, tmp_size, wav->fp);
            if(read_len < tmp_size){
                printf("error wav file\n");
                wav_close(&wav);
                fclose(dest_fp);
                return -1;
            }
            wav->format.compression_code = *(short *)buffer;
            wav->format.channels = *(short *)(buffer + 2);
            wav->format.samples_per_sec = *(int *)(buffer + 4);
            wav->format.avg_bytes_per_sec = *(int *)(buffer + 8);
            wav->format.block_align = *(short *)(buffer + 12);
            wav->format.bits_per_sample = *(short *)(buffer + 14);

            wav_write_header(wav,dest_fp);
        }
        else if(0 == strncasecmp("DATA", id_buffer, 4)){
            int rdlen,wtlen;
            int filllen = 0;
            memcpy(wav->data.id, id_buffer, 4);
            wav->data.size = tmp_size;
            offset += 8;
            wav->data_offset = offset;
            wav->data_size = wav->data.size;

            write_len = fwrite(wav->data.id, 1, 4, dest_fp);
            write_len = fwrite((char *)(&(wav->data.size)), 1, 4, dest_fp);

            while(filllen < wav->data_size){
                rdlen = fread(buffer, 1, 1024, wav->fp);
                if(rdlen == 0){
                    printf("read end\n");
                    wav_close(&wav);
                    fclose(dest_fp);
                    return 0;
                }
                filllen+=rdlen;
                wav_sample_amp((int16_t*)buffer,rdlen/2,factor);
                //printf("tell:%d.\n",ftell(wav->fp));
                //fseek(wav->fp, -rdlen, SEEK_CUR);
                wtlen = fwrite(buffer, 1, rdlen, dest_fp);
            }

            while(1){
                rdlen = fread(buffer, 1, 1024, wav->fp);
                if(rdlen == 0){
                    printf("read end\n");
                    wav_close(&wav);
                    fclose(dest_fp);
                    return 0;
                }

                wtlen = fwrite(buffer, 1, rdlen, dest_fp);
            }
            break;
        }
        else{
            printf("unhandled chunk: %s, size: %d\n", id_buffer, tmp_size);
            fseek(wav->fp, tmp_size, SEEK_CUR);
        }
        offset += 8 + tmp_size;
    }
    wav_close(&wav);
    fclose(dest_fp);
    return 0;
}

int wav_sample_amp(int16_t *sample,int len,float factor){
    if(factor > 1){
        return 0;
    }

    for(int i=0;i<len;i++){
        *(sample+i) = (int16_t)((*(sample+i)) * factor);
        
    }
    
    return 0;
}

int wav_getdb(double factor){
    return factor * 48 - 48;
}

///////////////////wav.c////////////////////
// int main(int argc, char **argv){
//     wav_t *wav = NULL;
//     if(argc <2){
//         printf("param error.\n");
//         exit(0);
//     }

//     int iret = wav_convert(argv[1],"test.wav",0.5);
//     if(iret < 0){
//         printf("convert error.\n");
//     }
//     return 1;
// }