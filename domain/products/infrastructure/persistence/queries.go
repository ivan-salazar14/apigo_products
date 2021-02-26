package persistence

const (

	// insertRole is a query that inserts a new row in the role table
	insertProduct = "INSERT INTO public.productos (producto_id,producto_nombre,producto_cantidad,producto_usercreacion,producto_fechcreacion,producto_usermodificacion,producto_fechamodificacion) VALUES ($1,$2,$3,$4,'NOW()',$5,'NOW()')"
)
