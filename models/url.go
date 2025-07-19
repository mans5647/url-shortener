package models


type NewFormUrl struct
{
	Id 				int				`json:"id" gorm:"primaryKey;autoIncrement;unique"`
	Code 			string			`json:"code"`
	ExpiringTime 	int				`json:"time"`
	RealUrl 		string			`json:"real_url"`
	ShortUrl 		string			`json:"short_url"`
}

type OldFormUrl struct
{
	Url 			string			`json:"url"`
	ExpiringTime 	int				`json:"time"`
}