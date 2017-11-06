/**
 * Created by VoidArtanis on 10/24/2017
 */

package models

type SimpleCRUD interface {
	Create() (bool,error)
	Update()( bool,error)
	Delete() (bool,error)

}
