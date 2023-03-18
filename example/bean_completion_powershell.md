---
title: "bean completion powershell"
description: bean completion powershell

---
## bean completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	bean completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
bean completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -d, --doit   do the thing
```

### SEE ALSO

* [bean completion](bean_completion/)	 - Generate the autocompletion script for the specified shell

