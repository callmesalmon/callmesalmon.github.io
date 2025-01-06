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

    "goarch/utils"
)

const TEMPL = `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title></title>
    </head>
    <body>
    
    </body>
</html>`

func main() {
    if len(os.Args) <= 2 {
        fmt.Println("USAGE: goarch <arcv> <dir>")
        return
    }
    archive := os.Args[1]
    dir := utils.Ls(os.Args[2])
    for _, file := range dir {
        err := os.WriteFile("archive/" + archive + "/" + file,
            []byte(utils.Templ(archive, file, false) + utils.Cat(os.Args[2] + "/" + file)), 0666)
        if err != nil {
            log.Fatalf("Failed to write to file \"%s\" :(", file)
        }
    }
    err := os.WriteFile("archive/" + archive + ".html",
        []byte(utils.Templ(archive, "", true) + TEMPL), 0666)
    if err != nil {
        log.Fatalf("Failed to write to file \"%s.html\" :(", archive)
    }
}
