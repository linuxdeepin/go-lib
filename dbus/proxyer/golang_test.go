package main

import "fmt"
import "testing"
import "os/exec"
import "time"

func TestGolang(t *testing.T) {
	if err := exec.Command("go", "install").Run(); err != nil {
		t.Fatal("install:" + err.Error())
	}

	var cmd *exec.Cmd
	cmd = exec.Command("proxyer", "-out", "tmp_golang__", "-target", "golang")
	if err := cmd.Start(); err != nil {
		t.Fatal("Proxyer:" + err.Error())
	}
	if err := cmd.Wait(); err != nil {
		t.Fatal("ProxyerWait:" + err.Error())
	}
	<-time.After(time.Millisecond * 50)

	/*if err := exec.Command("bash", "-c", "cd tmp_golang__ && go build").Run(); err != nil {*/
	if out, err := exec.Command("bash", "-c", "cd tmp_golang__ && ls && go build").CombinedOutput(); err != nil {
		fmt.Println(string(out))
		t.Fatal("Build:" + err.Error())
	}
	if err := exec.Command("rm", "-rf", "tmp_golang__").Run(); err != nil {
		t.Fatal("rm:" + err.Error())
	}
}
