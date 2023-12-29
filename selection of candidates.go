func Completion() *cobra.Command {
	// completionCmd outputs the completion script
	c := &cobra.Command{
		Use:   "completions",
		Short: "Generate tab completion scripts",


      $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  To load completions for each session, execute once:
  
  $ pylonsd completions zsh > "${fpath[1]}/_pylonsd"

  You will need to start a new shell for this setup to take effect.

fish:

  $ pylonsd completions fish | source}}


func DevCelCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cel-check [expresssion] [type] [[var_type, var_name, var_value]...]",
		Short: "Compiles and runs a cel expression with the provided variables and return type",
		Args:  cobra.MinimumNArgs(2),}}
