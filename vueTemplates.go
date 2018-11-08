package main

import (
	"bytes"
	"html/template"
)

const tpl = `<template>
	<div class="container">
		{{.Name}}
	</div>
</template>

<script>
export default {
	name: '{{.Name}}',
	data () {
		return {

		}
	}
}
</script>

<style lang="css" scoped>
</style>`

func writeTemplate(name string) string {
	t, err := template.New("webpage").Parse(tpl)
	check(err)
	data := struct {
		Name string
	}{
		Name: name}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	check(err)

	return buf.String()
}
