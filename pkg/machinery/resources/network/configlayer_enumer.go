// Code generated by "enumer -type=ConfigLayer,Operator -linecomment -text"; DO NOT EDIT.

package network

import (
	"fmt"
	"strings"
)

const _ConfigLayerName = "defaultcmdlineplatformoperatorconfiguration"

var _ConfigLayerIndex = [...]uint8{0, 7, 14, 22, 30, 43}

const _ConfigLayerLowerName = "defaultcmdlineplatformoperatorconfiguration"

func (i ConfigLayer) String() string {
	if i < 0 || i >= ConfigLayer(len(_ConfigLayerIndex)-1) {
		return fmt.Sprintf("ConfigLayer(%d)", i)
	}
	return _ConfigLayerName[_ConfigLayerIndex[i]:_ConfigLayerIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ConfigLayerNoOp() {
	var x [1]struct{}
	_ = x[ConfigDefault-(0)]
	_ = x[ConfigCmdline-(1)]
	_ = x[ConfigPlatform-(2)]
	_ = x[ConfigOperator-(3)]
	_ = x[ConfigMachineConfiguration-(4)]
}

var _ConfigLayerValues = []ConfigLayer{ConfigDefault, ConfigCmdline, ConfigPlatform, ConfigOperator, ConfigMachineConfiguration}

var _ConfigLayerNameToValueMap = map[string]ConfigLayer{
	_ConfigLayerName[0:7]:        ConfigDefault,
	_ConfigLayerLowerName[0:7]:   ConfigDefault,
	_ConfigLayerName[7:14]:       ConfigCmdline,
	_ConfigLayerLowerName[7:14]:  ConfigCmdline,
	_ConfigLayerName[14:22]:      ConfigPlatform,
	_ConfigLayerLowerName[14:22]: ConfigPlatform,
	_ConfigLayerName[22:30]:      ConfigOperator,
	_ConfigLayerLowerName[22:30]: ConfigOperator,
	_ConfigLayerName[30:43]:      ConfigMachineConfiguration,
	_ConfigLayerLowerName[30:43]: ConfigMachineConfiguration,
}

var _ConfigLayerNames = []string{
	_ConfigLayerName[0:7],
	_ConfigLayerName[7:14],
	_ConfigLayerName[14:22],
	_ConfigLayerName[22:30],
	_ConfigLayerName[30:43],
}

// ConfigLayerString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ConfigLayerString(s string) (ConfigLayer, error) {
	if val, ok := _ConfigLayerNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ConfigLayerNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ConfigLayer values", s)
}

// ConfigLayerValues returns all values of the enum
func ConfigLayerValues() []ConfigLayer {
	return _ConfigLayerValues
}

// ConfigLayerStrings returns a slice of all String values of the enum
func ConfigLayerStrings() []string {
	strs := make([]string, len(_ConfigLayerNames))
	copy(strs, _ConfigLayerNames)
	return strs
}

// IsAConfigLayer returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ConfigLayer) IsAConfigLayer() bool {
	for _, v := range _ConfigLayerValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for ConfigLayer
func (i ConfigLayer) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ConfigLayer
func (i *ConfigLayer) UnmarshalText(text []byte) error {
	var err error
	*i, err = ConfigLayerString(string(text))
	return err
}

const _OperatorName = "dhcp4dhcp6vip"

var _OperatorIndex = [...]uint8{0, 5, 10, 13}

const _OperatorLowerName = "dhcp4dhcp6vip"

func (i Operator) String() string {
	if i < 0 || i >= Operator(len(_OperatorIndex)-1) {
		return fmt.Sprintf("Operator(%d)", i)
	}
	return _OperatorName[_OperatorIndex[i]:_OperatorIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _OperatorNoOp() {
	var x [1]struct{}
	_ = x[OperatorDHCP4-(0)]
	_ = x[OperatorDHCP6-(1)]
	_ = x[OperatorVIP-(2)]
}

var _OperatorValues = []Operator{OperatorDHCP4, OperatorDHCP6, OperatorVIP}

var _OperatorNameToValueMap = map[string]Operator{
	_OperatorName[0:5]:        OperatorDHCP4,
	_OperatorLowerName[0:5]:   OperatorDHCP4,
	_OperatorName[5:10]:       OperatorDHCP6,
	_OperatorLowerName[5:10]:  OperatorDHCP6,
	_OperatorName[10:13]:      OperatorVIP,
	_OperatorLowerName[10:13]: OperatorVIP,
}

var _OperatorNames = []string{
	_OperatorName[0:5],
	_OperatorName[5:10],
	_OperatorName[10:13],
}

// OperatorString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func OperatorString(s string) (Operator, error) {
	if val, ok := _OperatorNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _OperatorNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Operator values", s)
}

// OperatorValues returns all values of the enum
func OperatorValues() []Operator {
	return _OperatorValues
}

// OperatorStrings returns a slice of all String values of the enum
func OperatorStrings() []string {
	strs := make([]string, len(_OperatorNames))
	copy(strs, _OperatorNames)
	return strs
}

// IsAOperator returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Operator) IsAOperator() bool {
	for _, v := range _OperatorValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for Operator
func (i Operator) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Operator
func (i *Operator) UnmarshalText(text []byte) error {
	var err error
	*i, err = OperatorString(string(text))
	return err
}
