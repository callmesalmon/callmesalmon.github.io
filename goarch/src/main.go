package main

/* * * * * * * * * * * * * * * * * * * * * * * * * *
 *                GOARCH TUTORIAL                  *
 * * * * * * * * * * * * * * * * * * * * * * * * * *
 * BUILDING                                        *
 *    For this, you need a Go compiler. Simply     *
 *    run the command ``./make.sh`` in the source  *
 *    directory (It's likely you need sudo         *
 *    permissions for this).                       *
 *                                                 *
 * PREREQUISITES                                   *
 *     To run goarch, you need a preexisting       *
 *     archive directory and invoke that           *
 *     directory as argv[1].                       *
 *                                                 *
 * RUNNING                                         *
 *     Invoke goarch like this:                    *
 *                                                 *
 *         goarch <archive> <source>               *
 *                                                 *
 *     Goarch is awfully simple, so you can't      *
 *     invoke with any specific flags.             *
 * * * * * * * * * * * * * * * * * * * * * * * * * */

import (
    "fmt"
    "log"
    "os"
)

func cat(file string) string {
    content, err := os.ReadFile(file)
    if err != nil {
        log.Fatalf("Couldn't read file \"%s\" :(", file)
    }
    return string(content)
}

func ls(dir string) []string {
    f, err := os.Open(dir)
    if err != nil {
        log.Fatalf("Couldn't open directory \"%s\" :(", dir)
    }
    files, err := f.Readdir(0)
    if err != nil {
        log.Fatalf("Couldn't walk directory \"%s\" :(", dir)
    }

    filenms := []string{}

    for _, v := range files {
        filenms = append(filenms, v.Name())
    }

    return filenms
}

/* It's hella ugly, ik */
func templ(arcv string, file string) string {
    return "---\n" + "permalink: /archive/" + arcv + "/" + file + "\n---\n"
}

func main() {
    if len(os.Args) <= 2 {
        fmt.Println("USAGE: archive <arcv> <dir>")
        return
    }
    archive := os.Args[1]
    dir := ls(os.Args[2])
    for _, file := range dir {
        err := os.WriteFile("archive/" + archive + "/" + file,
            []byte(templ(archive, file) + cat(os.Args[2] + "/" + file)), 0666)
        if err != nil {
            log.Fatalf("Failed to write to file \"%s\" :(", file)
        }
    }
}
