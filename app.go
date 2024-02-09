package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type publishment struct {
	file   string
	title  string
	replys int
}

var f_matches = []string{"<div class=comment id", "comentario svelte" /*, "tem-comentario", ... */}

func f_identify(f_file string) publishment {
	f_publishment := publishment{file: f_file, title: f_file, replys: 0}
	if strings.Contains(f_publishment.title, "/") {
		f_publishment.title = f_publishment.title[strings.LastIndex(f_publishment.title, "/")+1:]
	}
	f_bytes, f_err := os.ReadFile(f_file)
	if f_err == nil {
		f_publishment.replys = -1
	}
	f_text := string(f_bytes)
	for i := 0; i < len(f_matches); i++ {
		f_publishment.replys += strings.Count(f_text, f_matches[i])
	}
	return f_publishment
}

func f_join(f_routes []string) []publishment {
	var f_publishments []publishment
	for i := 0; i < len(f_routes); i++ {
		f_publishments = append(f_publishments, f_identify(f_routes[i]))
	}
	sort.Slice(f_publishments[:], func(i, j int) bool {
		return f_publishments[i].replys < f_publishments[j].replys
	})

	return f_publishments
}

func f_files(f_arguments []string) []string {
	var f_files []string
	for i := 0; i < len(f_arguments); i++ {
		f_name := f_arguments[i]
		f_folder, f_err := os.ReadDir(f_arguments[i])
		if f_err != nil {
			f_files = append(f_files, f_arguments[i])
		}
		if f_name[len(f_name)-1:] != "/" {
			f_name = f_name + "/"
		}
		for _, f_file := range f_folder {
			if f_file.IsDir() != true {
				f_files = append(f_files, f_name+f_file.Name())
			}
		}
	}
	return f_files

}

func f_between(f_del string, f_string string) []string {
	var f_files []string
	f_return := strings.Split(f_string, f_del)
	for i := 0; i < len(f_return); i++ {
		if len(f_return[i]) > 3 {
			f_files = append(f_files, f_return[i])
		}
	}
	return f_files
}

func f_routes(f_arguments []string) ([]string, time.Time) {
	if len(f_arguments) > 0 {
		return f_arguments, time.Now()
	}
	var f_subargss = []string{}
	var f_subargs string
	fmt.Println("\n[***] Emm... no ingresaste ninguna ruta, pon al menos un archivo o carpeta:")
	f_scanner := bufio.NewScanner(os.Stdin)
	for f_scanner.Scan() {
		f_subargs = f_scanner.Text()
		break
	}
	if strings.Count(f_subargs, "\"") > 1 {
		f_subargss = f_between("\"", f_subargs)
	}
	if strings.Count(f_subargs, "'") > 1 {
		f_subargss = f_between("'", f_subargs)
	}
	return f_subargss, time.Now()
}

func main() {
	fmt.Println("\n[***] Hola, aquí «respuestasantricas» en ejecución. ")
	f_routes, f_start := f_routes(os.Args[1:])
	f_files := f_files(f_routes)
	fmt.Println("\n[***] Ok... Se combrobarán", len(f_files), "archivos a ordenar.")
	f_join := f_join(f_files)
	for i := 0; i < len(f_join); i++ {
		fmt.Println(f_join[i].replys, " ---> ", f_join[i].title)
	}
	fmt.Println("\n[***] Comprobaciones completadas! En", time.Since(f_start), "salió eso.")
	fmt.Println("[***] Ahora estás viendo el total de respuestas en esas", len(f_join), "publicaciones.")
}
