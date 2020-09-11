#ifndef _WAV_H_
#define _WAV_H_
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef struct _wav_riff_t{
    char id[5]; //ID:"RIFF"
    int size; //file_len - 8
    char type[5]; //type:"WAVE"
}wav_riff_t;


typedef struct _wav_format_t{
    char id[5]; //ID:"fmt"
    int size;
    short compression_code;
    short channels;
    int samples_per_sec;
    int avg_bytes_per_sec;
    short block_align;
    short bits_per_sample;
}wav_format_t;


typedef struct _wav_fact_t{
    char id[5];
    int size;
}wav_fact_t;


typedef struct _wav_data_t{
    char id[5];
    int size;
}wav_data_t;


typedef struct _wav_t{
    FILE *fp;
    wav_riff_t riff;
    wav_format_t format;
    wav_fact_t fact;
    wav_data_t data;
    int file_size;
    int data_offset;
    int data_size;
}wav_t;
 

wav_t *wav_open(char *file_name);

//int strncasecmp(char *s1, char *s2, register int n);

void wav_close(wav_t **wav);

void wav_rewind(wav_t *wav);

int wav_over(wav_t *wav);

int wav_read_data(wav_t *wav, char *buffer, int buffer_size);

void wav_dump(wav_t *wav);

int wav_convert(char *file_name,char *dest_name,float precent);

#endif