package main

type Quota int8

const (
	GEN Quota = iota
	SC
	ST
	OBC
	PD
	DUNNO
)

var quotaMap = map[Quota]string{
	GEN: "Your All India Rank is",
	SC:  "Your Rank in the SC List is",
	ST:  "Your Rank in the ST List is",
	OBC: "Your Rank in the OBC(NCL) List",
	PD:  "Your Rank in the PD List is",
}

var regionMap = map[int][]string{
	101: {"panaji", "goa"},

	102: {"ahmedabad", "gujarat"},
	103: {"surat", "gujarat"},
	104: {"vadodara", "gujarat"},

	105: {"mumbai", "maharashtra"},
	106: {"nagpur", "maharashtra"},
	107: {"navi mumbai", "maharashtra"},
	108: {"pune", "maharashtra"},

	109: {"ajmer", "rajasthan"},
	110: {"jaipur", "rajasthan"},
	111: {"jodhpur", "rajasthan"},

	201: {"delhi east", "delhi"},
	202: {"delhi west", "delhi"},
	203: {"delhi north", "delhi"},
	204: {"delhi south", "delhi"},
	205: {"delhi central", "delhi"},

	206: {"faridabad", "haryana"},
	207: {"gurgaon", "haryana"},

	208: {"jammu", "jammu and kashmir"},

	209: {"indore", "madhya pradesh"},

	210: {"sikar", "rajasthan"},
	211: {"udaipur", "rajasthan"},

	212: {"aligarh", "uttar pradesh"},
	213: {"mathura", "uttar pradesh"},

	214: {"dubai", "uae"},

	301: {"itanagar", "arunachal pradesh"},

	302: {"guhawati", "assam"},
	303: {"johrat", "assam"},
	304: {"silchar", "assam"},

	305: {"gaya", "bihar"},
	306: {"katihar", "bihar"},
	307: {"muzaffarpur", "bihar"},
	308: {"patna", "bihar"},

	309: {"imphal", "manipur"},

	310: {"shillong", "meghalaya"},

	311: {"siliguri", "west bengal"},

	401: {"bhopal", "madhya pradesh"},
	402: {"gwalior", "madhya pradesh"},
	403: {"jabalpur", "madhya pradesh"},

	404: {"pantnagar", "uttarakhand"},

	405: {"agra", "uttar pradesh"},
	406: {"allahabad", "uttar pradesh"},
	407: {"gorakhpur", "uttar pradesh"},
	408: {"jhansi", "uttar pradesh"},
	409: {"kanpur", "uttar pradesh"},
	410: {"lucknow", "uttar pradesh"},

	501: {"port blair", "andaman and nicobar islands"},

	502: {"visakhapatnam", "andhra pradesh"},

	503: {"bhilai", "chattisgarh"},
	504: {"bilaspur", "chattisgarh"},
	505: {"raipur", "chattisgarh"},

	506: {"bokaro", "jharkhand"},
	507: {"dhanbad", "jharkhand"},
	508: {"jamshedpur", "jharkhand"},
	509: {"ranchi", "jharkhand"},

	510: {"bhubaneswar", "orissa"},
	511: {"rourkela", "orissa"},

	512: {"gangtok", "sikkim"},

	513: {"agartala", "tripura"},

	514: {"durgapur", "west bengal"},
	515: {"kharagpur", "west bengal"},
	516: {"kolkata north", "west bengal"},
	517: {"kolkata salt lake", "west bengal"},
	518: {"kolkata south", "west bengal"},
	519: {"malda", "west bengal"},

	601: {"hyderabad", "andhra pradesh"},
	602: {"nellore", "andhra pradesh"},
	603: {"vijaywada", "andhra pradesh"},
	604: {"warangal", "andhra pradesh"},

	605: {"balangore", "karnataka"},
	606: {"magalore", "karnataka"},

	607: {"kochi", "kerala"},
	608: {"kozhikode", "kerala"},
	609: {"thiruvananthapuram", "kerala"},

	610: {"puducherry", "puducherry"},

	611: {"chennai", "tamil nadu"},
	612: {"madurai", "tamil nadu"},

	701: {"chandigarh", "chandigarh"},

	702: {"kurukshetra", "haryana"},
	703: {"panipat", "haryana"},
	704: {"rohtak", "haryana"},

	705: {"palampur", "himachal pradesh"},
	706: {"shimla", "himachal pradesh"},

	707: {"amritsar", "punjab"},
	708: {"bathinda", "punjab"},
	709: {"jalandhar", "punjab"},
	710: {"ludhiana", "punjab"},
	711: {"patiala", "punjab"},

	712: {"dehradun", "uttarakhand"},
	713: {"roorkee", "uttarakhand"},

	714: {"bareilly", "uttar pradesh"},
	715: {"noida", "uttar pradesh"},
	716: {"ghaziabad", "uttar pradesh"},
	717: {"meerut", "uttar pradesh"},
	718: {"moradabad", "uttar pradesh"},
	719: {"varanasi", "uttar pradesh"},
}

type Student struct {
	Selected  bool
	Rollno    int
	Region    int
	Name      string
	Rank      int
	Q         Quota
	Plaintext string
}
