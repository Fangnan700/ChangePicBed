package model

type Config struct {
	InputDir    string      `mapstructure:"input_dir"`
	OutputDir   string      `mapstructure:"output_dir"`
	TempDir     string      `mapstructure:"temp_dir"`
	PicBed      string      `mapstructure:"pic_bed"`
	CosConfig   cosConfig   `mapstructure:"cos_config"`
	YuqueConfig yuqueConfig `mapstructure:"yuque_config"`
}

type cosConfig struct {
	BucketName string `mapstructure:"bucket_name"`
	BucketArea string `mapstructure:"bucket_area"`
	PicPath    string `mapstructure:"pic_path"`
	SecretID   string `mapstructure:"secret_id"`
	SecretKey  string `mapstructure:"secret_key"`
}

type yuqueConfig struct {
	YuqueSession string `mapstructure:"_yuque_session"`
	YuqueCToken  string `mapstructure:"yuque_ctoken"`
	ExportPath   string `mapstructure:"export_path"`
}
