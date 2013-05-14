package main

import (
    "os/exec"
    "os"
    "math/rand"
    "syscall"
)


const name_len = 7
var random_bytes = []byte("1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
var temp_bytes = make([]byte, name_len)
func random_name() string {
    for i:=0;i<name_len;i++ {
        temp_bytes[i] = random_bytes[rand.Int31n(62)]
    }
    return string(temp_bytes)
}
func process_code(code string, tp string) (string, error) {
    var filen string
    for true {
        filen = random_name()
        err := syscall.Access(outfile_dir+filen, syscall.F_OK)
        if err!=nil {
            break
        }
    }
    filename := filen+"."+tp
    infile, err := os.Create(infile_dir+filename)
    if err != nil {
        return  "", err
    }
    infile.WriteString(code)
    infile.Close()

    err = os.Chdir("highlight_py")
    if err != nil {
        //return  err
        return  "",err
    }
    defer os.Chdir("..")

    cmd := exec.Command("python","highlight.py",filename)
    err = cmd.Run()
    if err != nil {
        //return  err
        return  "",err
    }
    return filen, nil
}
