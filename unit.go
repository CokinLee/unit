package unit

import "fmt"

// Unit represents a unit
type Unit interface {
	// Float64 return base number
	Float64() float64
}

// Type 单位类型
type Type string

const (
	// UnitTypeNumber 数字 十、百、千、万...
	UnitTypeNumber Type = "Number"
	// UnitTypeAcceleration 加速度
	UnitTypeAcceleration = "Acceleration"
	// UnitTypeAmountOfSubstance 物质的量  摩尔
	UnitTypeAmountOfSubstance = "AmountOfSubstance"
	// UnitTypeAngle 角度和弧度
	UnitTypeAngle = "Angle"
	// UnitTypeArea 面积 平方米 公顷等
	UnitTypeArea = "Area"
	// UnitTypeDatarate 信息传输速率 Kb/s KB/s 等
	UnitTypeDatarate = "Datarate"
	// UnitTypeDatasize 数据大小
	UnitTypeDatasize = "Datasize"
	// UnitTypeDuration Duration 时间间隔 年月日时分秒
	UnitTypeDuration = "Duration"
	// UnitTypeElectricCurrent 电流
	UnitTypeElectricCurrent = "ElectricCurrent"
	// UnitTypeElectricalConductance 电导 siemens西门子(电导单位)
	UnitTypeElectricalConductance = "ElectricalConductance"
	// UnitTypeElectricalResistance 电阻
	UnitTypeElectricalResistance = "ElectricalResistance"
	// UnitTypeEnergy 能量 焦耳Joule / 瓦时watthour / Calorie卡路里(Gramcalorie) Kilocalorie千卡 Megacalorie 兆卡
	UnitTypeEnergy = "Energy"
	// UnitTypeForce 力 Newton(牛/牛顿) Dyne(达因) (KilogramForce/Kilopond)千克力
	UnitTypeForce = "Force"
	// UnitTypeFrequency 频率 Hertz(Hz赫兹) SI ...
	UnitTypeFrequency = "Frequency"
	// UnitTypeIlluminance (光) 照度 Lux(勒克斯（照明单位）)
	UnitTypeIlluminance = "Illuminance"
	// UnitTypeLength 长度 Meter 米SI... // US Inch Hand Foot Yard Link Rod Chain Furlong Mile // space LunarDistance AstronomicalUnit LightYear
	UnitTypeLength = "Length"
	// UnitTypeLuminousFlux 光通量、光束（流）  Lumen(流明)
	UnitTypeLuminousFlux = "LuminousFlux"
	// UnitTypeLuminousIntensity 发光强度、照度、光强 Candela(坎，坎德拉)
	UnitTypeLuminousIntensity = "LuminousIntensity"
	// UnitTypeMass 质量、重量
	UnitTypeMass = "Mass"
	// UnitTypePower 功率、电功率  瓦、瓦特(Watt)
	UnitTypePower = "Power"
	// UnitTypePressure 压力;压强;大气压;  帕(帕斯卡)Pascal
	UnitTypePressure = "Pressure"
	// UnitTypeSpeed 速度  MetersPerSecond 米每秒（m/s） km/h
	UnitTypeSpeed = "Speed"
	// UnitTypeTemperature 温度 kelvin（开尔文/摄氏度)
	UnitTypeTemperature = "Temperature"
	// UnitTypeVoltage 电压;伏特数  Volt（伏特）
	UnitTypeVoltage = "Voltage"
	// UnitTypeVolume 体积;容量 Liter(升) CubicMeter(立方米)
	UnitTypeVolume = "Volume"
	// UnitTypeCurrency 货币
	UnitTypeCurrency = "Currency"
)

// NewUnit create a new unit
func NewUnit(f float64, unitType Type) (Unit, error) {
	switch unitType {
	case UnitTypeNumber:
		return Number(f), nil
	case UnitTypeAcceleration:
		return Acceleration(f), nil
	case UnitTypeAmountOfSubstance:
		return AmountOfSubstance(f), nil
	case UnitTypeAngle:
		return Angle(f), nil
	case UnitTypeArea:
		return Area(f), nil
	case UnitTypeDatarate:
		return Datarate(f), nil
	case UnitTypeDatasize:
		return Datasize(f), nil
	case UnitTypeDuration:
		return Duration(f), nil
	case UnitTypeElectricCurrent:
		return ElectricCurrent(f), nil
	case UnitTypeElectricalConductance:
		return ElectricalConductance(f), nil
	case UnitTypeElectricalResistance:
		return ElectricalResistance(f), nil
	case UnitTypeEnergy:
		return Energy(f), nil
	case UnitTypeForce:
		return Force(f), nil
	case UnitTypeFrequency:
		return Frequency(f), nil
	case UnitTypeIlluminance:
		return Illuminance(f), nil
	case UnitTypeLength:
		return Length(f), nil
	case UnitTypeLuminousFlux:
		return LuminousFlux(f), nil
	case UnitTypeLuminousIntensity:
		return LuminousIntensity(f), nil
	case UnitTypeMass:
		return Mass(f), nil
	case UnitTypePower:
		return Power(f), nil
	case UnitTypePressure:
		return Pressure(f), nil
	case UnitTypeSpeed:
		return Speed(f), nil
	case UnitTypeTemperature:
		return Temperature(f), nil
	case UnitTypeVoltage:
		return Voltage(f), nil
	case UnitTypeVolume:
		return Volume(f), nil
	case UnitTypeCurrency:
		return Currency(f), nil
	default:
		return 0 * Base, fmt.Errorf("unknown unit type %s", unitType)
	}
}
