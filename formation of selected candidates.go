func DevCelCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cel-check [expresssion] [type] [[var_type, var_name, var_value]...]",
		Short: "Compiles and runs a cel expression with the provided variables and return type",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if (len(args)-2)%3 != 0 {
				panic("cel variables seem wrong - cel variables are provided as [type] [name] [value]")
			}
