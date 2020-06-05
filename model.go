package emailrep

// SearchResponse is a reputation response from the api
type SearchResponse struct {
	Email      string  `json:"email"`      // email address queried
	Reputation string  `json:"reputation"` // high/medium/low/none
	Suspicious bool    `json:"suspicious"` // total number of positive and negative sources of reputation. note that these may not all be direct references to the email address, but can include reputation sources for the domain or other related information
	References int64   `json:"references"` // whether the email address should be treated as suspicious or risky
	Details    Details `json:"details"`
}

// Details of the email from emailrep.io
type Details struct {
	Blacklisted             bool          `json:"blacklisted"`                // the email is believed to be malicious or spammy
	MaliciousActivity       bool          `json:"malicious_activity"`         // the email has exhibited malicious behavior (e.g. phishing or fraud)
	MaliciousActivityRecent bool          `json:"malicious_activity_recent"`  // malicious behavior in the last 90 days (e.g. in the case of temporal account takeovers)
	CredentialsLeaked       bool          `json:"credentials_leaked"`         // credentials were leaked at some point in time (e.g. a data breach, pastebin, dark web, etc.)
	CredentialsLeakedRecent bool          `json:"credentials_leaked_recent"`  // credentials were leaked in the last 90 days
	DataBreach              bool          `json:"data_breach"`                // the email was in a data breach at some point in time
	FirstSeen               string        `json:"first_seen"`                 // the first date the email was observed in a breach, credential leak, or exhibiting malicious or spammy behavior (‘never’ if never seen)
	LastSeen                string        `json:"last_seen"`                  // the last date the email was observed in a breach, credential leak, or exhibiting malicious or spammy behavior (‘never’ if never seen)
	DomainExists            bool          `json:"domain_exists"`              // valid domain
	DomainReputation        string        `json:"domain_reputation"`          // high/medium/low/n/a (n/a if the domain is a free_provider, disposable, or doesn’t exist)
	NewDomain               bool          `json:"new_domain"`                 // the domain was created within the last year
	DaysSinceDomainCreation int64         `json:"days_since_domain_creation"` // days since the domain was created
	SuspiciousTLD           bool          `json:"suspicious_tld"`             // suspicious tld
	Spam                    bool          `json:"spam"`                       // the email has exhibited spammy behavior (e.g. spam traps, login form abuse)
	FreeProvider            bool          `json:"free_provider"`              // the email uses a free email provider
	Disposable              bool          `json:"disposable"`                 // the email uses a temporary/disposable service
	Deliverable             bool          `json:"deliverable"`                // deliverable
	AcceptAll               bool          `json:"accept_all"`                 // whether the mail server has a default accept all policy. some mail servers return inconsistent responses, so we may default to an accept_all for those to be safe
	ValidMX                 bool          `json:"valid_mx"`                   // has an MX record
	Spoofable               bool          `json:"spoofable"`                  // email address can be spoofed (e.g. not a strict SPF policy or DMARC is not enforced)
	SPFStrict               bool          `json:"spf_strict"`                 // sufficiently strict SPF record to prevent spoofing
	DmarcEnforced           bool          `json:"dmarc_enforced"`             // DMARC is configured correctly and enforced
	Profiles                []interface{} `json:"profiles"`                   // online profiles used by the email
	Summary                 string        // Human readable summary only returned if `summary=true`
}
