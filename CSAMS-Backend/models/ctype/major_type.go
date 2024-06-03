package ctype

import "encoding/json"

type Major int

const (
	EnglishLanguage                                    Major = 1  // 英语
	FinancialManagement                                Major = 2  // 财务管理
	Accounting                                         Major = 3  // 会计学
	HumanResourceManagement                            Major = 4  // 人力资源管理
	Marketing                                          Major = 5  // 市场营销
	JapaneseLanguage                                   Major = 6  // 日语
	NetworkAndNewMedia                                 Major = 7  // 网络与新媒体
	PoliticalScienceAndPublicAdministration            Major = 8  // 政治学与行政学
	EngineeringMechanics                               Major = 9  // 工程力学
	ArtificialIntelligence                             Major = 10 // 人工智能
	IntelligentMiningEngineering                       Major = 11 // 智能采矿工程
	IntelligentManufacturingEngineering                Major = 12 // 智能制造工程
	EnvironmentalEngineering                           Major = 13 // 环境工程
	ChemicalEngineeringAndTechnology                   Major = 14 // 化学工程与工艺
	UrbanUndergroundSpaceEngineering                   Major = 15 // 城市地下空间工程
	InternetOfThingsEngineering                        Major = 16 // 物联网工程
	IndustrialDesign                                   Major = 17 // 工业设计
	DigitalMediaTechnology                             Major = 18 // 数字媒体技术
	CommunicationsEngineering                          Major = 19 // 通信工程
	CivilEngineering                                   Major = 20 // 土木工程
	ComputerScienceAndTechnology                       Major = 21 // 计算机科学与技术
	MathematicsAndAppliedMathematics                   Major = 22 // 数学与应用数学
	EngineeringCost                                    Major = 23 // 工程造价
	MechanicalDesignManufacturingAndAutomation         Major = 24 // 机械设计制造及其自动化
	IntelligentConstruction                            Major = 25 // 智能建造
	ElectricalEngineeringAndAutomation                 Major = 26 // 电气工程及其自动化
	PreventiveMedicine                                 Major = 27 // 预防医学
	GeographicInformationEngineering                   Major = 28 // 地理空间信息工程
	GeologicalEngineering                              Major = 29 // 地质工程
	Pharmacy                                           Major = 30 // 药学
	InformationSecurity                                Major = 31 // 信息安全
	SafetyEngineering                                  Major = 32 // 安全工程
	SurveyingEngineering                               Major = 33 // 测绘工程
	SoftwareEngineering                                Major = 34 // 软件工程
	AppliedChemistry                                   Major = 35 // 应用化学
	FireEngineering                                    Major = 36 // 消防工程
	ResourceExplorationEngineering                     Major = 37 // 资源勘查工程
	OptoelectronicInformationScienceAndEngineering     Major = 38 // 光电信息科学与工程
	PharmaceuticalEngineering                          Major = 39 // 制药工程
	InformationAndComputingScience                     Major = 40 // 信息与计算科学
	InorganicNonmetallicMaterialsEngineering           Major = 41 // 无机非金属材料工程
	Automation                                         Major = 42 // 自动化
	ProcessEquipmentAndControlEngineering              Major = 43 // 过程装备与控制工程
	HydrologyAndWaterResourcesEngineering              Major = 44 // 水文与水资源工程
	RoboticsEngineering                                Major = 45 // 机器人工程
	IntelligentScienceAndTechnology                    Major = 46 // 智能科学与技术
	Finance                                            Major = 47 // 金融学
	ClinicalMedicine                                   Major = 48 // 临床医学
	DataScienceAndBigDataTechnology                    Major = 49 // 数据科学与大数据技术
	MechatronicsEngineering                            Major = 50 // 机械电子工程
	AppliedStatistics                                  Major = 51 // 应用统计学
	MeasurementControlTechnologyAndInstrument          Major = 52 // 测控技术与仪器
	MiningEngineering                                  Major = 53 // 采矿工程
	IntegratedCircuitDesignAndIntegratedSystems        Major = 54 // 集成电路设计与集成系统
	ECommerce                                          Major = 55 // 电子商务
	Architecture                                       Major = 56 // 建筑学
	AmmunitionEngineeringAndExplosivesTechnology       Major = 57 // 弹药工程与爆炸技术
	PolymerMaterialsAndEngineering                     Major = 58 // 高分子材料与工程
	BlockchainEngineering                              Major = 59 // 区块链工程
	RemoteSensingScienceAndTechnology                  Major = 60 // 遥感科学与技术
	BuildingEnvironmentAndEnergyApplicationEngineering Major = 61 // 建筑环境与能源应用工程
	TransportationEngineering                          Major = 62 // 交通工程
	IntelligentMaterialsAndStructures                  Major = 63 // 智能材料与结构
	NavigationEngineering                              Major = 64 // 导航工程
	ElectronicInformationEngineering                   Major = 65 // 电子信息工程
	AppliedPhysics                                     Major = 66 // 应用物理学
	EmergencyTechnologyAndManagement                   Major = 67 // 应急技术与管理
	MedicalLaboratoryTechnology                        Major = 68 // 医学检验技术
	EnergyChemicalEngineering                          Major = 69 // 能源化学工程
	ExplorationTechnologyAndEngineering                Major = 70 // 勘查技术与工程
	VehicleEngineering                                 Major = 71 // 车辆工程
)

func (s Major) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Major) String() string {
	var str string
	switch s {
	case EnglishLanguage:
		str = "英语"
	case FinancialManagement:
		str = "财务管理"
	case Accounting:
		str = "会计学"
	case HumanResourceManagement:
		str = "人力资源管理"
	case Marketing:
		str = "市场营销"
	case JapaneseLanguage:
		str = "日语"
	case NetworkAndNewMedia:
		str = "网络与新媒体"
	case PoliticalScienceAndPublicAdministration:
		str = "政治学与行政学"
	case EngineeringMechanics:
		str = "工程力学"
	case ArtificialIntelligence:
		str = "人工智能"
	case IntelligentMiningEngineering:
		str = "智能采矿工程"
	case IntelligentManufacturingEngineering:
		str = "智能制造工程"
	case EnvironmentalEngineering:
		str = "环境工程"
	case ChemicalEngineeringAndTechnology:
		str = "化学工程与工艺"
	case UrbanUndergroundSpaceEngineering:
		str = "城市地下空间工程"
	case InternetOfThingsEngineering:
		str = "物联网工程"
	case IndustrialDesign:
		str = "工业设计"
	case DigitalMediaTechnology:
		str = "数字媒体技术"
	case CommunicationsEngineering:
		str = "通信工程"
	case CivilEngineering:
		str = "土木工程"
	case ComputerScienceAndTechnology:
		str = "计算机科学与技术"
	case MathematicsAndAppliedMathematics:
		str = "数学与应用数学"
	case EngineeringCost:
		str = "工程造价"
	case MechanicalDesignManufacturingAndAutomation:
		str = "机械设计制造及其自动化"
	case IntelligentConstruction:
		str = "智能建造"
	case ElectricalEngineeringAndAutomation:
		str = "电气工程及其自动化"
	case PreventiveMedicine:
		str = "预防医学"
	case GeographicInformationEngineering:
		str = "地理空间信息工程"
	case GeologicalEngineering:
		str = "地质工程"
	case Pharmacy:
		str = "药学"
	case InformationSecurity:
		str = "信息安全"
	case SafetyEngineering:
		str = "安全工程"
	case SurveyingEngineering:
		str = "测绘工程"
	case SoftwareEngineering:
		str = "软件工程"
	case AppliedChemistry:
		str = "应用化学"
	case FireEngineering:
		str = "消防工程"
	case ResourceExplorationEngineering:
		str = "资源勘查工程"
	case OptoelectronicInformationScienceAndEngineering:
		str = "光电信息科学与工程"
	case PharmaceuticalEngineering:
		str = "制药工程"
	case InformationAndComputingScience:
		str = "信息与计算科学"
	case InorganicNonmetallicMaterialsEngineering:
		str = "无机非金属材料工程"
	case Automation:
		str = "自动化"
	case ProcessEquipmentAndControlEngineering:
		str = "过程装备与控制工程"
	case HydrologyAndWaterResourcesEngineering:
		str = "水文与水资源工程"
	case RoboticsEngineering:
		str = "机器人工程"
	case IntelligentScienceAndTechnology:
		str = "智能科学与技术"
	case Finance:
		str = "金融学"
	case ClinicalMedicine:
		str = "临床医学"
	case DataScienceAndBigDataTechnology:
		str = "数据科学与大数据技术"
	case MechatronicsEngineering:
		str = "机械电子工程"
	case AppliedStatistics:
		str = "应用统计学"
	case MeasurementControlTechnologyAndInstrument:
		str = "测控技术与仪器"
	case MiningEngineering:
		str = "采矿工程"
	case IntegratedCircuitDesignAndIntegratedSystems:
		str = "集成电路设计与集成系统"
	case ECommerce:
		str = "电子商务"
	case Architecture:
		str = "建筑学"
	case AmmunitionEngineeringAndExplosivesTechnology:
		str = "弹药工程与爆炸技术"
	case PolymerMaterialsAndEngineering:
		str = "高分子材料与工程"
	case BlockchainEngineering:
		str = "区块链工程"
	case RemoteSensingScienceAndTechnology:
		str = "遥感科学与技术"
	case BuildingEnvironmentAndEnergyApplicationEngineering:
		str = "建筑环境与能源应用工程"
	case TransportationEngineering:
		str = "交通工程"
	case IntelligentMaterialsAndStructures:
		str = "智能材料与结构"
	case NavigationEngineering:
		str = "导航工程"
	case ElectronicInformationEngineering:
		str = "电子信息工程"
	case AppliedPhysics:
		str = "应用物理学"
	case EmergencyTechnologyAndManagement:
		str = "应急技术与管理"
	case MedicalLaboratoryTechnology:
		str = "医学检验技术"
	case EnergyChemicalEngineering:
		str = "能源化学工程"
	case ExplorationTechnologyAndEngineering:
		str = "勘查技术与工程"
	case VehicleEngineering:
		str = "车辆工程"
	default:
		str = "其他"
	}
	return str
}
