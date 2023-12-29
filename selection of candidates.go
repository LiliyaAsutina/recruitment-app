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

// selection of the best candidates
func DevCelCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cel-check [expresssion] [type] [[var_type, var_name, var_value]...]",
		Short: "Compiles and runs a cel expression with the provided variables and return type",
		Args:  cobra.MinimumNArgs(2),}}

			case "fish":
				err := cmd.Root().GenFishCompletion(os.Stdout, true)
				if err != nil {
					panic(err)

// passing a special selection test
	const (
	cookbookExtension = ".plc"
	recipeExtension   = ".plr"
	moduleExtension   = ".pdt"
	includeDirective  = "#include "
)
					
				}
			println("Result:")
			switch returnType {
			case "long":
				r, err := ec.EvalInt64(args[0])
				if err != nil {
					panic(err)
				}}
