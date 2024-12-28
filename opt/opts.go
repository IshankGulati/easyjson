package opt

//go:generate sed -i "s/\\+build none/generated by gotemplate/" optional/opt.go
//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Int(int)
//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Uint(uint)

//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Int8(int8)
//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Int16(int16)
//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Int32(int32)
//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Int64(int64)

//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Uint8(uint8)
//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Uint16(uint16)
//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Uint32(uint32)
//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Uint64(uint64)

//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Float32(float32)
//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Float64(float64)

//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" Bool(bool)
//go:generate gotemplate "github.com/IshankGulati/easyjson/opt/optional" String(string)
//go:generate sed -i "s/generated by gotemplate/+build none/" optional/opt.go
