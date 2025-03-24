package {{ .PackageName }}

func InsertOne(one *{{ .StructName }}) error{
    // TODO 插入逻辑
    return nil
}

func QueryOne() (*{{ .StructName }}, error){


    return &{{ .StructName }}{},nil
}
