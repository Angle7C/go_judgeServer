package config

type JudgeConfig struct {
	Work         string `yaml:"Work"`
	JudgePath    string `yaml:"JudgePath"`
	CompilerPath string `yaml:"CompilerPath"`
	ComparePath  string `yaml:"ComparePath"`
}
