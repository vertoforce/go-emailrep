package emailrep

// SearchResponse is a reputation response from the api
type SearchResponse struct {
	Email      string  `json:"email"`
	Reputation string  `json:"reputation"`
	Suspicious bool    `json:"suspicious"`
	References int64   `json:"references"`
	Details    Details `json:"details"`
}

type Details struct {
	Blacklisted             bool          `json:"blacklisted"`
	MaliciousActivity       bool          `json:"malicious_activity"`
	MaliciousActivityRecent bool          `json:"malicious_activity_recent"`
	CredentialsLeaked       bool          `json:"credentials_leaked"`
	CredentialsLeakedRecent bool          `json:"credentials_leaked_recent"`
	DataBreach              bool          `json:"data_breach"`
	FirstSeen               string        `json:"first_seen"`
	LastSeen                string        `json:"last_seen"`
	DomainExists            bool          `json:"domain_exists"`
	DomainReputation        string        `json:"domain_reputation"`
	NewDomain               bool          `json:"new_domain"`
	DaysSinceDomainCreation int64         `json:"days_since_domain_creation"`
	SuspiciousTLD           bool          `json:"suspicious_tld"`
	Spam                    bool          `json:"spam"`
	FreeProvider            bool          `json:"free_provider"`
	Disposable              bool          `json:"disposable"`
	Deliverable             bool          `json:"deliverable"`
	AcceptAll               bool          `json:"accept_all"`
	ValidMX                 bool          `json:"valid_mx"`
	Spoofable               bool          `json:"spoofable"`
	SPFStrict               bool          `json:"spf_strict"`
	DmarcEnforced           bool          `json:"dmarc_enforced"`
	Profiles                []interface{} `json:"profiles"`
}
