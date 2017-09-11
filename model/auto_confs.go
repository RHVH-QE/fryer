package model

// CommonParams store all project wide params
type CommonParams struct {
	ID                string // primary key
	BuildsServerURL   string `storm:"index"`
	CurrentIPPort     []string
	LogURL            string
	CobblerAPI        string
	CobblerCredential []string
}

// Host contain host information
type Host struct {
	ShortName  string `storm:"id"`
	BeakerName string `storm:"index"`
	NicName    string
	Mac        string
}

// KickstartScript is
type KickstartScript struct {
	ScriptName string `storm:"id"`
	Content    string
}

// AutoTestTiers is
type AutoTestTiers struct {
	ID            int // primary key
	DebugTier     int
	AnacondaTier1 int
	AnacondaTier2 int
	KsTier1       int
	KsTier2       int
	UpgradeTier1  int
	UpgradeTier2  int
	VdsmTier      int
}
