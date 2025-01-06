package utils

import (
    "os"
    "log"
)

func Cat(file string) string {
    content, err := os.ReadFile(file)
    if err != nil {
        log.Fatalf("Couldn't read file \"%s\" :(", file)
    }
    return string(content)
}

func Ls(dir string) []string {
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
func Templ(arcv string, file string, root bool) string {
    if (!root) {
        return "---\npermalink: /archive/" + arcv + "/" + file + "\n---\n\n"
    }
    return "---\npermalink: /archive/" + arcv + "\n---\n\n"
}
