package swagger

//go:generate rm -rf ../../internal/generated
//go:generate mkdir -p ../../internal/generated
//go:generate swagger generate server --quiet --target ../../internal/generated --name api --spec swagger.yml --exclude-main
