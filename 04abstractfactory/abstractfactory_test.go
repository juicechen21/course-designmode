package abstractfactory

import "testing"

func getMainAndDetail(fa DAOFactory)  {
	fa.CreateOrderMainDAO().SaveOrderMain()
	fa.CreateOrderDetailDAO().SaveOrderDetail()
}
func TestRdbFactory(t *testing.T) {
	var fa DAOFactory
	fa = &RDBDAOFactory{}
	getMainAndDetail(fa)
}

func TestXmlFactory(t *testing.T) {
	var fa DAOFactory
	fa = &XMLDAOFactory{}
	getMainAndDetail(fa)
}
