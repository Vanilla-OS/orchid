package cmdr

/* License: GPLv3
Authors:
    Mirko Brombin <send@mirko.pm>
    Pietro di Caprio <pietro@fabricators.ltd>
Copyright: 2023
Description: Orchid is a cli helper for Vanilla OS projects
*/

import (
    "fmt"
    "log"
    "os"
    "path"
    "path/filepath"
    "strings"
    "time"

    "github.com/spf13/cobra"
    "github.com/spf13/cobra/doc"
)

func NewDocsCommand(a *App) *Command {
    c := &Command{}
    cmd := &cobra.Command{
        Use:                   "docs",
        Short:                 "Generates documentation for the cli application in the current directory.",
        SilenceUsage:          true,
        DisableFlagsInUseLine: true,
        Hidden:                true,
        Args:                  cobra.NoArgs,
        RunE: func(cmd *cobra.Command, args []string) error {
            filePrepender := func(filename string) string {
                name := filepath.Base(filename)
                base := strings.TrimSuffix(name, path.Ext(name))
                title := capitalizeFirst(base)
                creationDate := getFileCreationDate(filename).Format("2006-01-02")
                return fmt.Sprintf(fmTemplate, title, title, creationDate)
            }

            // Generate documentation
            err := doc.GenMarkdownTreeCustom(cmd.Root(), "./", filePrepender, func(name string) string {
                base := strings.TrimSuffix(name, path.Ext(name))
                return strings.ToLower(base) + "/"
            })

            if err != nil {
                log.Fatal(err)
            }

            return nil
        },
    }
    c.Command = cmd
    return c
}

func capitalizeFirst(s string) string {
    if len(s) == 0 {
        return s
    }
    return strings.ToUpper(s[:1]) + s[1:]
}

func getFileCreationDate(filename string) time.Time {
    fileInfo, err := os.Stat(filename)
    if err != nil {
        log.Fatal(err)
    }
    return fileInfo.ModTime()
}

const fmTemplate = `---
Title: %s Manpage
Description: Manpage for the %s utility.
PublicationDate: %s
Authors: Contributors of Vanilla OS
---

`
