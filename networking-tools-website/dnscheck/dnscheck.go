package dnscheck

import (
	"net"

	"github.com/miekg/dns"

	"github.com/gokhangokcen1/subnet-backend/models"
)

const dnsServer = "8.8.8.8:53"

func CheckAllRecords(host string) models.DNSFullResult {
	result := models.DNSFullResult{Domain: host}

	ips, _ := net.LookupIP(host)

	for _, ip := range ips {
		if ip.To4() != nil {
			result.A = append(result.A, ip.String())
		} else {
			result.AAAA = append(result.AAAA, ip.String())
		}
	}

	if cname, err := net.LookupCNAME(host); err == nil {
		result.CNAME = cname
	}

	if mxs, err := net.LookupMX(host); err == nil {
		for _, mx := range mxs {
			result.MX = append(result.MX, models.MXRecord{
				Preference: mx.Pref,
				Host:       mx.Host,
			})
		}
	}

	if nss, err := net.LookupNS(host); err == nil {
		for _, ns := range nss {
			result.NS = append(result.NS, ns.Host)
		}
	}

	if txts, err := net.LookupTXT(host); err == nil {
		result.TXT = txts
	}

	for _, ip := range ips {
		if ip.To4() == nil {
			continue
		}
		if ptrs, err := net.LookupAddr(ip.String()); err == nil {
			result.PTR = ptrs
		}
		break
	}

	if _, srvs, err := net.LookupSRV("sip", "tcp", host); err == nil {
		for _, srv := range srvs {
			result.SRV = append(result.SRV, models.SRVRecord{
				Priority: srv.Priority,
				Weight:   srv.Weight,
				Port:     srv.Port,
				Target:   srv.Target,
			})
		}
	}

	result.SOA = querySOA(host)
	result.DNSKEY = queryDNSKEY(host)
	result.DS = queryDS(host)
	result.CAA = queryCAA(host)

	return result
}

func querySOA(domain string) *models.SOARecord {
	answers := exchangeQuery(domain, dns.TypeSOA)
	for _, ans := range answers {
		if soa, ok := ans.(*dns.SOA); ok {
			return &models.SOARecord{
				NS:      soa.Ns,
				Mbox:    soa.Mbox,
				Serial:  soa.Serial,
				Refresh: soa.Refresh,
				Retry:   soa.Retry,
				Expire:  soa.Expire,
				Minimum: soa.Minttl,
			}
		}
	}
	return nil
}

func queryDNSKEY(domain string) []models.DNSKEYRecord {
	var result []models.DNSKEYRecord
	answers := exchangeQuery(domain, dns.TypeDNSKEY)
	for _, ans := range answers {
		if key, ok := ans.(*dns.DNSKEY); ok {
			result = append(result, models.DNSKEYRecord{
				Flags:     key.Flags,
				Protocol:  key.Protocol,
				Algorithm: key.Algorithm,
				PublicKey: key.PublicKey,
			})
		}
	}
	return result
}

func queryDS(domain string) []models.DSRecord {
	var result []models.DSRecord
	answers := exchangeQuery(domain, dns.TypeDS)
	for _, ans := range answers {
		if ds, ok := ans.(*dns.DS); ok {
			result = append(result, models.DSRecord{
				KeyTag:     ds.KeyTag,
				Algorithm:  ds.Algorithm,
				DigestType: ds.DigestType,
				Digest:     ds.Digest,
			})
		}
	}
	return result
}

func queryCAA(domain string) []models.CAARecord {
	var result []models.CAARecord
	answers := exchangeQuery(domain, dns.TypeCAA)
	for _, ans := range answers {
		if caa, ok := ans.(*dns.CAA); ok {
			result = append(result, models.CAARecord{
				Flag:  caa.Flag,
				Tag:   caa.Tag,
				Value: caa.Value,
			})
		}
	}
	return result
}

func exchangeQuery(domain string, recordType uint16) []dns.RR {
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(domain), recordType)

	client := new(dns.Client)
	resp, _, err := client.Exchange(msg, dnsServer)
	if err != nil || resp == nil {
		return nil
	}
	return resp.Answer
}
