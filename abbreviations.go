package scs

// The abbreviations contains a list of words that has
// a specific format in Camel and Pascal Case styles.
var abbreviations = map[string]string{
	"5g":      "5G",      // 5th generation
	"ack":     "ACK",     // Acknowledgement
	"acl":     "ACL",     // Access Control List
	"adfs":    "ADFS",    // Active Directory Federation Service
	"adsl":    "ADSL",    // Asymmetric digital subscriber line
	"aes":     "AES",     // Advanced Encryption System
	"afaik":   "AFAIK",   // As far as I Know
	"aka":     "AKA",     // Also known as
	"alu":     "ALU",     // Arithmetic Logic Units
	"ansi":    "ANSI",    // American National Standards Institute
	"ap":      "AP",      // Access Point
	"api":     "API",     // Application Programming Interface
	"arp":     "ARP",     // Address Resolution Protocol
	"asap":    "ASAP",    // As soon as possible
	"ascii":   "ASCII",   // American Standard Code for Information Interchange
	"asn":     "ASN",     // Autonomous System Number
	"atm":     "ATM",     // Asynchronous Transfer Mode
	"au":      "AU",      // Activation Unit
	"avx":     "AVX",     // Advanced Vector Extensions
	"bf":      "BF",      // Boyfriend
	"bgp":     "BGP",     // Border Gateway Protocol
	"bi":      "BI",      // Business Intelligence
	"bios":    "BIOS",    // Basic Input Output System
	"bpo":     "BPO",     // Business Process Outsourcing
	"brb":     "BRB",     // Be right back
	"bss":     "BSS",     // Basic service set (Wi-Fi)
	"btw":     "BTW",     // By the way
	"byod":    "BYOD",    // Bring Your Own Device
	"cat":     "CAT",     // Category
	"cbsp":    "CBSP",    // Cloud-based Security Providers
	"ccitt":   "CCITT",   // Standards organization
	"cdn":     "CDN",     // Content Delivery Network
	"chap":    "CHAP",    // Challenge-Handshake Authentication Protocol (PPP)
	"cidr":    "CIDR",    // Classless Inter-Domain Routing
	"cir":     "CIR",     // Committed Information Rate (Frame Relay)
	"cli":     "CLI",     // Command Line Interface
	"cmos":    "CMOS",    // Complementary metal-oxide semiconductor
	"cname":   "CNAME",   // Canonical Name
	"cors":    "CORS",    // Cross-Origin Resource Sharing
	"cpe":     "CPE",     // Customer premises equipment
	"cpu":     "CPU",     // Central processing Unit
	"crc":     "CRC",     // Cyclic redundancy check
	"crt":     "CRT",     // Cathode Ray Tube
	"csp":     "CSP",     // Content Security Policy
	"csrf":    "CSRF",    // Cross-Site Request Forgery
	"css":     "CSS",     // Cascading Style Sheets
	"cta":     "CTA",     // Call to action
	"ctr":     "CTR",     // Click-through rate
	"cvss":    "CVSS",    // Common Vulnerability Scoring System
	"cwe":     "CWE",     // Common Weakness Enumeration
	"d2d":     "D2D",     // Device to Device
	"dam":     "DAM",     // Database activity monitoring
	"dast":    "DAST",    // Dynamic Application Security Testing
	"dc":      "DC",      // Data Center
	"dce":     "DCE",     // Data communications equipment
	"ddos":    "DDoS",    // Distributed Denial Of Services
	"dec":     "DEC",     // Digital Equipment Corporation
	"des":     "DES",     // Data Encryption Standard
	"dhcp":    "DHCP",    // Dynamic Host Configuration Protocol
	"dimm":    "DIMM",    // Dual In-line Memory Module
	"diy":     "DIY",     // Do it yourself
	"dkim":    "DKIM",    // Domain Keys Identified Mail
	"dm":      "DM",      // Direct Message
	"dmarc":   "DMARC",   // Domain-based Msg. Auth. Reporting and Conformance
	"dmi":     "DMI",     // Desktop Management Interface
	"dns":     "DNS",     // Domain Name System
	"dnssec":  "DNSSEC",  // Domain Name System Security Extensions
	"dom":     "DOM",     // Document Object Model
	"dos":     "DOS",     // Denial of Services
	"dram":    "DRAM",    // Dynamic random-access memory
	"drs":     "DRS",     // Distributed Resource Scheduler
	"dsa":     "DSA",     // Digital Signature Algorithm
	"dsl":     "DSL",     // Digital subscriber line
	"dslam":   "DSLAM",   // Digital subscriber line access multiplexer
	"dt":      "DT",      // Directory Traversal
	"dte":     "DTE",     // Data Terminal Equipment
	"dvd":     "DVD",     // Digital Versatile Disc
	"dw":      "DW",      // Data Warehouse
	"daas":    "DaaS",    // Desktop-as-a-Service
	"eha":     "EHA",     // Ethernet Hardware Address (MAC address)
	"eia":     "EIA",     // Electronics Industry Alliance
	"eigrp":   "EIGRP",   // Enhanced Interior Gateway Routing Protocol
	"eod":     "EOD",     // End of day
	"eof":     "EOF",     // End Of File
	"eol":     "EOL",     // End Of Life
	"erp":     "ERP",     // Enterprise Resource Planning
	"ess":     "ESS",     // Extended service set (Wi-Fi group)
	"esx":     "ESX",     // Elastic Sky X
	"faq":     "FAQ",     // Frequently asked question
	"fb":      "FB",      // Facebook
	"fcc":     "FCC",     // Federal Communications Commission (US)
	"fcs":     "FCS",     // Frame check sequence (Ethernet)
	"fddi":    "FDDI",    // Fiber Distributed Data Interface
	"fifo":    "FIFO",    // First In First Out
	"foss":    "FOSS",    // Free and open-source software
	"ftp":     "FTP",     // File Transfer Protocol
	"fud":     "FUD",     // Fully Undetectable
	"fwiw":    "FWIW",    // For what it’s worth
	"gbic":    "GBIC",    // Gigabit interface converter
	"gepof":   "GEPOF",   // Gigabit Ethernet (over) Plastic Optical Fiber
	"gf":      "GF",      // Girlfriend
	"gui":     "GUI",     // Graphical User Interface
	"ha":      "HA",      // High Availability
	"hac":     "HAC",     // High-availability Clusters
	"har":     "HAR",     // HTTP Archive
	"hdlc":    "HDLC",    // High-level Data Link Control
	"hmu":     "HMU",     // Hit me up
	"html":    "HTML",    // Hypertext Markup Language
	"http":    "HTTP",    // HyperText Transport Protocol
	"https":   "HTTPS",   // HyperText Transport Protocol Secure
	"i0":      "I0",      // Input & Output
	"iam":     "IAM",     // Identity & Access Management
	"iana":    "IANA",    // Internet Assigned Number Authority
	"icmp":    "ICMP",    // Internet Control Message Protocol
	"icymi":   "ICYMI",   // In case you missed it
	"icaas":   "ICaaS",   // Integration Capability as a Service
	"id":      "ID",      // Identifier
	"ide":     "IDE",     // Integrated Development Environment
	"idf":     "IDF",     // Intermediate distribution frame
	"idk":     "IDK",     // I don’t know
	"ids":     "IDS",     // Intrusion Detection System
	"ie":      "IE",      // Internet Explorer
	"ieee":    "IEEE",    // Institute for Electrical and Electronic Engineers
	"ietf":    "IETF",    // Internet Engineering Task Force
	"ifttt":   "IFTTT",   // If This Then That
	"ily":     "ILY",     // I love you
	"imap":    "IMAP",    // Internet Message Access Protocol
	"imho":    "IMHO",    // In my humble opinion
	"imo":     "IMO",     // In my opinion
	"iops":    "IOPS",    // Input/output Operations Per Second
	"ip":      "IP",      // Internet Protocol
	"ips":     "IPS",     // Intrusion prevention system
	"ipsec":   "IPSec",   // Internet Protocol Security
	"isis":    "ISIS",    // Intermediate System to Intermediate System
	"isdn":    "ISDN",    // Integrated Services Digital Network
	"isp":     "ISP",     // Internet Service Provider
	"it":      "IT",      // Information Technology
	"itut":    "ITUT",    // International Telecommunications Union
	"iot":     "IoT",     // Internet Of Things
	"jce":     "JCE",     // Java Cryptography Extension
	"jdk":     "JDK",     // Java Development Kit
	"js":      "JS",      // JavaScript
	"json":    "JSON",    // JavaScript Object Notation
	"kvm":     "KVM",     // Kernel-based Virtual Machine
	"lacp":    "LACP",    // Link Aggregation Control Protocol
	"lan":     "LAN",     // Local Area Network
	"lapb":    "LAPB",    // Link Access Procedure, Balanced (x.25)
	"lapf":    "LAPF",    // Link-access procedure for frame relay
	"lb":      "LB",      // Load Balancer
	"lfi":     "LFI",     // Local File Inclusion
	"llc":     "LLC",     // Logical link control
	"llt":     "LLT",     // Low Latency Transport
	"lmgtfy":  "LMGTFY",  // Let me Google that for you
	"lmk":     "LMK",     // Let me know
	"lol":     "LOL",     // Laugh out loud
	"mac":     "MAC",     // Media Access Control
	"mam":     "MAM",     // Media access management
	"man":     "MAN",     // Metropolitan area network
	"mc":      "MC",      // Multiple choice
	"mcm":     "MCM",     // Man crush Monday
	"mdf":     "MDF",     // Main distribution frame
	"mfa":     "MFA",     // Multi-Factor Authentication
	"mib":     "MIB",     // Management information base (SNMP)
	"mime":    "MIME",    // Multipurpose Internet Mail Extensions
	"mitm":    "MITM",    // Man in the Middle Attack
	"mms":     "MMS",     // Multimedia messaging service
	"mpls":    "MPLS",    // Multi-Protocol Label Switching
	"mpp":     "MPP",     // Massive Parallel Processing
	"mstsc":   "MSTSC",   // Microsoft Terminal Service Client
	"mtu":     "MTU",     // Maximum Transmission Unit
	"mx":      "MX",      // Mail Exchange
	"mbps":    "Mbps",    // Megabits per second
	"moca":    "MoCA",    // Multimedia over Coax Alliance
	"na":      "NA",      // Not applicable or not available
	"nac":     "NAC",     // Network access control
	"nack":    "NACK",    // Negative ACKnowledgement
	"nat":     "NAT",     // Network Address Translation
	"nbma":    "NBMA",    // Non-Broadcast Multiple Access
	"nda":     "NDA",     // Nondisclosure Agreement
	"nfs":     "NFS",     // Network File System
	"nic":     "NIC",     // Network Interface Card
	"np":      "NP",      // No problem
	"nrz":     "NRZ",     // Non-return-to-zero
	"nrzi":    "NRZI",    // Non-return to zero inverted
	"ns":      "NS",      // Name Server
	"nvm":     "NVM",     // Nevermind
	"nvram":   "NVRAM",   // Non-volatile RAM
	"ola":     "OLA",     // Operational-level Agreement
	"ooo":     "OOO",     // Out of office
	"ootd":    "OOTD",    // Outfit of the day
	"os":      "OS",      // Operating System
	"osci":    "OSCI",    // OS Command Injection
	"osi":     "OSI",     // Open System Interconnect
	"ospf":    "OSPF",    // Open Shortest Path First
	"oss":     "OSS",     // Open Source Software
	"oui":     "OUI",     // Organization Unique Identifier
	"owasp":   "OWASP",   // Open Web Application Security Project
	"p2v":     "P2V",     // Physical to Virtual
	"pap":     "PAP",     // Password authentication protocol
	"pat":     "PAT",     // Port address translation
	"pc":      "PC",      // Personal computer (host)
	"pcm":     "PCM",     // Pulse-code modulation
	"pdf":     "PDF",     // Portable Document Format
	"pdu":     "PDU",     // Protocol data unit
	"php":     "PHP",     // Hypertext Preprocessor
	"pim":     "PIM",     // Personal information manager
	"pm":      "PM",      // Private Message
	"pop":     "POP",     // Post Office Protocol
	"pop3":    "POP3",    // Post Office Protocol, version 3
	"post":    "POST",    // Power-on self test
	"pots":    "POTS",    // Plain old telephone service
	"ppp":     "PPP",     // Point-to-point Protocol
	"pptp":    "PPTP",    // Point-to-Point Tunneling Protocol
	"psu":     "PSU",     // Power Supply Unit
	"pt":      "PT",      // Path Traversal
	"ptt":     "PTT",     // Public Telephone and Telegraph
	"putin":   "KHUYLO",  // La lala la-la, la-lala ...
	"pvst":    "PVST",    // Per-VLAN Spanning Tree
	"paas":    "PaaS",    // Platform-as-a-Service
	"qemu":    "QEMU",    // Quick Emulator
	"qotd":    "QOTD",    // Quote of the day
	"quic":    "QUIC",    // Quick UDP Internet Connections
	"qos":     "QoS",     // Quality of Service
	"rad":     "RAD",     // Rapid Application Development
	"radius":  "RADIUS",  // Remote Authentication Dial-In User Service
	"raid":    "RAID",    // Redundant Array of Independent Disk
	"ram":     "RAM",     // Random-access Memory
	"rarp":    "RARP",    // Reverse ARP
	"rat":     "RAT",     // Remote Administration Tool
	"rcp":     "RCP",     // Royal College of Physicians
	"rcs":     "RCS",     // Rich communication services
	"rdp":     "RDP",     // Remote Desktop Protocol
	"rfc":     "RFC",     // Request for Comments
	"rfi":     "RFI",     // Remote File Inclusion
	"rimm":    "RIMM",    // Rambus In-line Memory Module
	"rip":     "RIP",     // Routing Information Protocol
	"risc":    "RISC",    // Reduced Instruction Set Computer
	"rll":     "RLL",     // Run-Length Limited
	"rn":      "RN",      // Right now
	"rofl":    "ROFL",    // Rolling on floor laughing
	"roi":     "ROI",     // Return on investment
	"rom":     "ROM",     // Read Only Memory
	"rss":     "RSS",     // Really Simple Syndication
	"rstp":    "RSTP",    // Rapid Spanning Tree Protocol
	"rtp":     "RTP",     // Real-time Transport Protocol
	"rum":     "RUM",     // Real-user Measurement
	"shttp":   "SHTTP",   // Secure HTTP (rarely used)
	"saml":    "SAML",    // Security Assertion Markup Language
	"san":     "SAN",     // Storage Area Network
	"sast":    "SAST",    // Static Application Security Testing
	"scd":     "SCD",     // Source Code Disclosure
	"scm":     "SCM",     // Supply Chain Management
	"sdk":     "SDK",     // Software Development Kit
	"sdlc":    "SDLC",    // Synchronous Data Link Control
	"sdn":     "SDN",     // Software Defined Networking
	"se":      "SE",      // Social Engineering
	"sfd":     "SFD",     // Start-of-frame delimiter
	"sfp":     "SFP",     // Small form-factor pluggable
	"sftp":    "SFTP",    // SSH File Transfer Protocol
	"simm":    "SIMM",    // Single In-line Memory Module
	"skid":    "SKid",    // Script Kiddie
	"sla":     "SLA",     // Service Level Agreement
	"slarp":   "SLARP",   // Serial Line ARP (Address Resolution Protocol)
	"slip":    "SLIP",    // Serial Line Internet Protocol
	"sms":     "SMS",     // Short message service
	"smtp":    "SMTP",    // Simple Mail Transfer Protocol
	"sna":     "SNA",     // Systems Network Architecture
	"snap":    "SNAP",    // SubNet Access Protocol
	"snmp":    "SNMP",    // Simple Network Management Protocol
	"soa":     "SOA",     // Start of Authority
	"soap":    "SOAP",    // Simple Object Access Protocol
	"sof":     "SOF",     // Start of frame
	"spf":     "SPF",     // Sender Policy Framework
	"sqli":    "SQLi",    // SQL Injection
	"sram":    "SRAM",    // Static random access memory
	"ssd":     "SSD",     // Solid-state Drive
	"sse":     "SSE",     // Server Side Encryption
	"ssh":     "SSH",     // Secure Shell
	"ssid":    "SSID",    // Service set identifier (Wi-Fi)
	"ssl":     "SSL",     // Secure Socket Layer
	"sso":     "SSO",     // Single Sign-On
	"stfu":    "STFU",    // Shut the *swear word!* up
	"stp":     "STP",     // Spanning Tree Protocol
	"sts":     "STS",     // Security Token Service
	"saas":    "SaaS",    // Software-as-a-Service
	"tba":     "TBA",     // To be announced
	"tbd":     "TBD",     // To be decided
	"tbf":     "TBF",     // To be frank
	"tbh":     "TBH",     // To be honest
	"tcp":     "TCP",     // Transmission Control Protocol
	"tdm":     "TDM",     // Time-division multiplexing
	"tftp":    "TFTP",    // Trivial File Transfer Protocol
	"tgif":    "TGIF",    // Thank goodness it’s Friday
	"tia":     "TIA",     // Telecommunications Industry Alliance
	"tld":     "TLD",     // Top-level domain
	"tldr":    "TLDR",    // Too long, didn’t read
	"tls":     "TLS",     // Transport Layer Security
	"tofu":    "TOFU",    // Trust On First Use
	"tps":     "TPS",     // Transaction Per Second
	"tpu":     "TPU",     // Tensor Processing Unit
	"ttf":     "TTF",     // TrueType Font
	"ttl":     "TTL",     // Time To Live
	"tty":     "TTY",     // Teletype
	"ttyl":    "TTYL",    // Talk to you later
	"udp":     "UDP",     // User Datagram Protocol
	"ugc":     "UGC",     // User-generated content
	"ups":     "UPS",     // Uninterruptible Power Supply
	"uri":     "URI",     // Uniform Resource Identifier
	"url":     "URL",     // Uniform Resource Locator
	"usb":     "USB",     // Universal Serial Bus
	"utp":     "UTP",     // Unshielded twisted pair
	"vbs":     "VBS",     // Visual Basic Script
	"vc":      "VC",      // Virtual circuit
	"vcs":     "VCS",     // Version Control Systems
	"vdi":     "VDI",     // Virtual Desktop Infrastructure
	"vg":      "VG",      // Volume Group
	"vip":     "VIP",     // Virtual IP
	"vlan":    "VLAN",    // Virtual Local Area Network
	"vlsm":    "VLSM",    // Variable-length subnet masking
	"vm":      "VM",      // Virtual Machine
	"vpn":     "VPN",     // Virtual Private Network
	"vsm":     "VSM",     // Virtual Supervisor Module
	"vtl":     "VTL",     // Virtual Tape Library
	"vxlan":   "VXLAN",   // Virtual Extensible Local Area Network
	"w3c":     "W3C",     // World Wide Web Consortium
	"waf":     "WAF",     // Web Application Firewall
	"wan":     "WAN",     // Wide Area Network
	"wap":     "WAP",     // Web Application Protection
	"wcw":     "WCW",     // Woman crush Wednesday
	"wep":     "WEP",     // Wired Equivalent Privacy
	"wins":    "WINS",    // Windows Internet Name Service
	"wlan":    "WLAN",    // Wireless Local Area Network
	"wlc":     "WLC",     // Wireless LAN Controller
	"wmi":     "WMI",     // Windows Management Instrumentation
	"wpa":     "WPA",     // Wi-Fi Protected Access
	"wpan":    "WPAN",    // Wireless Personal Area Network
	"www":     "WWW",     // World Wide Web
	"wysiwyg": "WYSIWYG", // What You See Is What You Get
	"wifi":    "WiFi",    // Wireless Fidelity
	"xhtml":   "XHTML",   // Extensible Hypertext Markup Language
	"xml":     "XML",     // Extensible Markup Language
	"xsd":     "XSD",     // XML Schema Definition
	"xss":     "XSS",     // Cross-site Scripting
	"iscsi":   "iSCSI",   // Internet Small Computer Storage Interface
}
