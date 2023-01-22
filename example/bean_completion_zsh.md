---
title: "bean completion zsh"
description: bean completion zsh

---
## bean completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(bean completion zsh); compdef _bean bean

To load completions for every new session, execute once:

#### Linux:

	bean completion zsh > "${fpath[1]}/_bean"

#### macOS:

	bean completion zsh > $(brew --prefix)/share/zsh/site-functions/_bean

You will need to start a new shell for this setup to take effect.


```
bean completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -d, --doit   do the thing
```

### SEE ALSO

* [bean completion](bean_completion/)	 - Generate the autocompletion script for the specified shell

