package models

type SslCheckRequest struct {
	Website string `json:"website"`
	Port    int    `json:"port"`
}

// ---- General Information bölümü (tek nesne, sertifika başına değil) ----
type GeneralInformation struct {
	ResolvesTo       string // google.com
	ExpirationDate   string // Sep 14, 2026
	VendorSigned     bool   // Yes
	HostnameMatches  bool   // Doesn't match
	KeyLength        int    // 256
	ServerType       string // NA
	RevocationStatus string // "Good", "Revoked", "Unknown", "NA"
}

// ---- Issued For bölümü (tek nesne - sadece LEAF sertifikanın Subject'i) ----
type IssuedFor struct {
	CommonName       string
	SAN              []string
	Organization     string
	OrganizationUnit string
	Country          string
	State            string
	Locality         string
	Address          string
}

// ---- Issued By bölümü (tek nesne - sadece LEAF sertifikanın Issuer'ı) ----
type IssuedBy struct {
	CommonName       string
	Organization     string
	OrganizationUnit string
	Country          string
	State            string
	Locality         string
}

// ---- Chain Details + Advanced bölümü (BİRDEN FAZLA - zincirdeki her sertifika için bir tane) ----
type ChainCert struct {
	Issuer             string
	CommonName         string
	Organization       string
	Issued             string
	Expires            string
	SerialNumber       string
	SignatureAlgorithm string
	FingerprintSHA1    string
	FingerprintMD5     string
}

type ChainDetails struct {
	Certs []ChainCert
}

// ---- Hepsini birleştiren üst rapor ----
type SSLReport struct {
	General GeneralInformation
	For     IssuedFor
	By      IssuedBy
	Chain   ChainDetails
}
