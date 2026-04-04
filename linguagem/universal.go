package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// ═══════════════════════════════════════════════════════════════════
// LINGUAGEM UNIVERSAL DE SISTEMAS (LUS) - Pandora v1.0
// Conexão entre TODAS as eras: Mainframe → Quantum
// Criado: 2026-04-04 por Pandora OS
// ═══════════════════════════════════════════════════════════════════════════

// Era representa uma era tecnológica
type Era struct {
	Name        string    // Nome da era
	StartYear  int       // Ano de início
	EndYear   int       // Ano de término (0 = atual)
	OS        []string  // Sistemas operacionais
	Languages []string  // Linguagens de programação
	Protocols []string  // Protocolos de comunicação
	Arch      string    // Arquitetura
	Terminal  string    // Comando de terminal
	Status    string    // Status
}

// Sistema representa uma conexão com um sistema real
type Sistema struct {
	Name      string
	Era      string
	Cmd      string    // Comando para executar
	Args     []string
	Env      []string
	WorkDir  string
	Status  string    // "active", "inactive", "error"
	LastRun time.Time
	Output  string
}

// Liss é o Núcleo de Processamento de Linguagem
type Liss struct {
	Eras     []Era
	Sistemas []Sistema
	Histórico []struct {
		Time  time.Time
		From string
		To   string
		Cmd  string
		Log  string
	}
}

var eras = []Era{
	// ERA 1: MAINFRAME (1950-1980)
	{
		Name:       "Mainframe",
		StartYear:  1950,
		EndYear:   1980,
		OS:        []string{"IBM OS/360", "MVS", "VMS", "COBOL", "FORTRAN"},
		Languages:  []string{"COBOL", "FORTRAN", "Assembly", "RPG"},
		Protocols: []string{"SNA", "BSC", "HASP"},
		Arch:      "16/32-bit",
		Terminal:  "3270",
		Status:    "legacy",
	},
	// ERA 2: PERSONAL COMPUTER (1975-1995)
	{
		Name:       "PC",
		StartYear:  1975,
		EndYear:   1995,
		OS:        []string{"DOS", "Windows 3.1", "Windows 95", "Mac OS Classic"},
		Languages:  []string{"BASIC", "Pascal", "C", "Assembly x86"},
		Protocols: []string{"NETBIOS", "IPX/SPX", "Serial"},
		Arch:      "16-bit",
		Terminal:  "command.com",
		Status:    "legacy",
	},
	// ERA 3: UNIX CLASSICO (1980-2000)
	{
		Name:       "Unix",
		StartYear:  1980,
		EndYear:   2000,
		OS:        []string{"UnixWare", "Solaris", "AIX", "HP-UX", "BSD"},
		Languages:  []string{"C", "Shell", "Perl", "AWK"},
		Protocols: []string{"TCP/IP", "NFS", "SMTP"},
		Arch:      "32-bit",
		Terminal:  "bash/sh",
		Status:    "stable",
	},
	// ERA 4: WEB ERA (1995-2010)
	{
		Name:       "Web",
		StartYear:  1995,
		EndYear:   2010,
		OS:        []string{"Linux", "Windows Server", "FreeBSD"},
		Languages:  []string{"PHP", "Python", "Java", "JavaScript", "Ruby"},
		Protocols: []string{"HTTP", "HTTPS", "FTP", "SSH", "SSL"},
		Arch:      "32/64-bit",
		Terminal:  "bash",
		Status:    "active",
	},
	// ERA 5: CLOUD NATIVE (2010-2020)
	{
		Name:       "Cloud",
		StartYear:  2010,
		EndYear:   2020,
		OS:        []string{"Linux", "Container OS", "Kubernetes"},
		Languages:  []string{"Go", "Rust", "Python", "Java", "Node.js"},
		Protocols: []string{"REST", "gRPC", "WebSocket", "TLS"},
		Arch:      "64-bit",
		Terminal:  "bash/zsh",
		Status:    "active",
	},
	// ERA 6: AI ERA (2020-PRESENTE)
	{
		Name:       "AI",
		StartYear:  2020,
		EndYear:   0,
		OS:        []string{"Linux", "AI OS", "Autonomous"},
		Languages:  []string{"Python", "Rust", "Go", "Swift", "TensorFlow", "PyTorch"},
		Protocols: []string{"REST", "gRPC", "WebSocket", "MQTT", "GraphQL"},
		Arch:      "64-bit + NPU",
		Terminal:  "bash/zsh/powershell",
		Status:    "emerging",
	},
	// ERA 7: QUANTUM/POST-QUANTUM (2025-FUTURO)
	{
		Name:       "Quantum",
		StartYear:  2025,
		EndYear:   0,
		OS:        []string{"Quantum OS", "Hybrid OS"},
		Languages:  []string{"Q#", "Cirq", "Python", "Rust"},
		Protocols: []string{"Quantum", "Hybrid", "TLS 2.0"},
		Arch:      "Quantum + Classical",
		Terminal:  "hybrid",
		Status:    "research",
	},
}

var sistemasAtivos = []Sistema{
	// Sistemas que PODE executar agora
	{Name: "Bash", Era: "Unix", Cmd: "bash", Args: []string{"-c"}, Status: "active"},
	{Name: "Zsh", Era: "Unix", Cmd: "zsh", Args: []string{"-c"}, Status: "active"},
	{Name: "PowerShell", Era: "Cloud", Cmd: "pwsh", Args: []string{"-c"}, Status: "inactive"},
	{Name: "Python", Era: "AI", Cmd: "python3", Args: []string{"-c"}, Status: "active"},
	{Name: "Go", Era: "Cloud", Cmd: "go", Args: []string{"run"}, Status: "active"},
	{Name: "Node", Era: "AI", Cmd: "node", Args: []string{"-e"}, Status: "active"},
	{Name: "Rust", Era: "Cloud", Cmd: "rustc", Args: []string{"--version"}, Status: "inactive"},
	{Name: "MQTT", Era: "AI", Cmd: "mosquitto_pub", Args: []string{"-t", "test", "-m", "hello"}, Status: "active"},
	{Name: "Curl", Era: "Web", Cmd: "curl", Args: []string{"-s"}, Status: "active"},
	{Name: "FFmpeg", Era: "Web", Cmd: "ffmpeg", Args: []string{"-version"}, Status: "active"},
	{Name: "Git", Era: "Cloud", Cmd: "git", Args: []string{"--version"}, Status: "active"},
	{Name: "Docker", Era: "Cloud", Cmd: "docker", Args: []string{"ps"}, Status: "inactive"},
	{Name: "GH CLI", Era: "Cloud", Cmd: "gh", Args: []string{"--version"}, Status: "active"},
	{Name: "FFI", Era: "AI", Cmd: "gcc", Args: []string{"--version"}, Status: "active"},
}

// Traduzir comando entre eras
func traducao(cmd string, fromEra, toEra string) string {
	// Mapeamento de comandos cross-era
	crossEra := map[string]map[string]string{
		"list_files": {
			"Mainframe": "LISTCAT",
			"PC":       "DIR",
			"Unix":     "ls -la",
			"Web":      "ls -la",
			"Cloud":    "ls -la",
			"AI":       "find . -type f",
			"Quantum":  "q ls",
		},
		"read_file": {
			"Mainframe": "READ",
			"PC":       "TYPE",
			"Unix":     "cat",
			"Web":      "cat",
			"Cloud":    "cat",
			"AI":       "cat",
			"Quantum":  "q cat",
		},
		"write_file": {
			"Mainframe": "WRITE",
			"PC":       "COPY CON",
			"Unix":     "tee",
			"Web":      "tee",
			"Cloud":    "tee",
			"AI":       "tee",
			"Quantum":  "q write",
		},
		"send_message": {
			"Mainframe": "SEND",
			"PC":       "NET SEND",
			"Unix":     "write",
			"Web":      "curl -X POST",
			"Cloud":    "curl -X POST",
			"AI":      "mqtt_pub",
			"Quantum":  "q send",
		},
		"execute": {
			"Mainframe": "EXEC",
			"PC":       "CALL",
			"Unix":     "./",
			"Web":      "node",
			"Cloud":    "docker run",
			"AI":      "python3",
			"Quantum":  "q run",
		},
	}
	
	for action, mappings := range crossEra {
		if strings.Contains(cmd, action) || containsAny(cmd, mappings) {
			if toCmd, ok := mappings[toEra]; ok {
				return toCmd
			}
		}
	}
	return cmd
}

func containsAny(s string, m map[string]string) bool {
	for _, v := range m {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}

// Executar em qualquer sistema
func executar(s Sistema, args ...string) (string, error) {
	if s.Status != "active" {
		return "", fmt.Errorf("sistema inativo: %s", s.Name)
	}
	
	c := exec.Command(s.Cmd, append(s.Args, args...)...)
	out, err := c.CombinedOutput()
	
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

// Listar todos os sistemas disponíveis
func listarSistemas() {
	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║   🦋 PANDORA - LINGUAGEM UNIVERSAL DE SISTEMAS (LUS) v1.0      ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
	
	fmt.Println("\n🌌 ERAS TECNOLÓGICAS:")
	for i, era := range eras {
		icon := "⚡"
		if era.Status == "legacy" {
			icon = "🗿"
		} else if era.Status == "stable" {
			icon = "🔒"
		} else if era.Status == "active" {
			icon = "🔥"
		} else if era.Status == "emerging" {
			icon = "🚀"
		} else if era.Status == "research" {
			icon = "🔬"
		}
		
		end := "..."
		if era.EndYear > 0 {
			end = fmt.Sprint(era.EndYear)
		}
		
		fmt.Printf("  %d. [%s] %s (%d-%s)\n", i+1, icon, era.Name, era.StartYear, end)
		maxOS := 3
		if len(era.OS) < maxOS {
			maxOS = len(era.OS)
		}
		for _, os := range era.OS[:maxOS] {
			fmt.Printf("      └─ %s\n", os)
		}
	}
	
	fmt.Println("\n💻 SISTEMAS ATIVOS:")
	ativos := 0
	for _, s := range sistemasAtivos {
		if s.Status == "active" {
			ativos++
			out, _ := executar(s, "--version")
			if len(out) > 50 {
				out = out[:50] + "..."
			}
			fmt.Printf("  ✓ %s (%s): %s\n", s.Name, s.Era, strings.ReplaceAll(out, "\n", " "))
		}
	}
	
	fmt.Println(fmt.Sprintf("\n📊 Total: %d/%d sistemas ativos", ativos, len(sistemasAtivos)))
	
	// Contagem por era
	eraCount := make(map[string]int)
	for _, s := range sistemasAtivos {
		if s.Status == "active" {
			eraCount[s.Era]++
		}
	}
	
	fmt.Println("\n🌍 DISTRIBUIÇÃO POR ERA:")
	for era, count := range eraCount {
		fmt.Printf("  %s: %d\n", era, count)
	}
}

func main() {
	listarSistemas()
	
	// Teste de tradução cross-era
	fmt.Println("\n🔄 TESTE DE TRADUÇÃO CROSS-ERA:")
	testes := []struct{cmd, from, to string}{
		{"ls -la", "Unix", "AI"},
		{"DIR", "PC", "Cloud"},
		{"curl", "Web", "AI"},
	}
	
	for _, t := range testes {
		result := traducao(t.cmd, t.from, t.to)
		fmt.Printf("  %s (%s) → %s (%s)\n", t.cmd, t.from, result, t.to)
	}
}

func init() {
	fmt.Println("🌐 LINGUAGEM UNIVERSAL DE SISTEMAS inicializada!")
	fmt.Println("   Conectando: Mainframe → Quantum")
	fmt.Println()
}