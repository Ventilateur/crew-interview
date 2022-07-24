package utils

func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func PtrString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
