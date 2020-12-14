package config

var (
	ZenURL            = "https://bankcbw.zendesk.com"
	CreateUserURL     = "https://bankcbwsupport.zendesk.com/api/v2/users/create_or_update.json"
	CreateManyuserURL = "https://bankcbw.zendesk.com/api/v2/users/create_or_update_many.json"
	GetExIDURL        = "https://bankcbwsupport.zendesk.com/api/v2/search.json?query=type:user+external_id:" //98989
	CreateTicketURL   = "https://bankcbwsupport.zendesk.com/api/v2/tickets.json"
	ShowTicketURL     = "https://bankcbw.zendesk.com/api/v2/tickets/"           //{id}.json
	SearchURL         = "https://bankcbw.zendesk.com/api/v2/search.json?query=" //{search_string}
	UserName          = "bakkarit@gmail.com"
	EndUserName       = "mechnaveen464@example.org/token:QhiVyLOaF9MrrtTiuvqAzorcMGSWHSC7d4kgsmgr"
	Password          = "Test@1234"
	MemberID          = 360031603334
	ApplicationID     = 360031601834
)

/*		search_string
query=3245227													The ticket with the id 3245227
query=Greenbriar												Any record with the word Greenbriar
query=type:user "Jane Doe"										User records with the exact string "Jane Doe"
query=type:ticket status:open									Open tickets
query=type:organization created<2015-05-01						Organizations created before May 1, 2015
query=status<solved requester:user@domain.com type:ticket		Unsolved tickets requested by user@domain.com
query=type:user tags:premium_support							Users tagged with premium_support
query=created>2012-07-17 type:ticket organization:"MD Photo"	Tickets created in the MD Photo org after July 17, 2012
oauth : 48a7f3dfd72ed8f60e29c3260a34fb3a5e1f61432d64dbada32f584a26cfa8cd
https://{subdomain}.zendesk.com/oauth/authorizations/new?response_type=code&redirect_uri={your_redirect_url}&client_id={your_unique_identifier}&scope=read%20write
https://bankcbw.zendesk.com/api/v2/users/400033187393/tickets/requested.json
*/
