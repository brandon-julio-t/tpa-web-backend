package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	_ "github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/factories"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/graph/resolvers"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	graphqlEndpoint = "/graphql"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print(err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowOrigins:     []string{"http://localhost:4200", "https://tpa-web-br20-2.netlify.app"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(new(middlewares.GinProviderMiddleware).Create())
	r.Use(new(middlewares.AuthProviderMiddleware).Create())

	r.GET("/", func(context *gin.Context) {
		playground.Handler("GraphQL playground", graphqlEndpoint).ServeHTTP(context.Writer, context.Request)
	})

	gql := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{},
	}))

	gql.Use(extension.AutomaticPersistedQuery{
		Cache: factories.NewCache(os.Getenv("REDIS_URL"), 24*time.Hour),
	})

	r.POST(graphqlEndpoint, func(context *gin.Context) {
		gql.ServeHTTP(context.Writer, context.Request)
	})

	runMigration()
	runSeed()

	log.Fatal(r.Run())
}

func runSeed() {
	if err := facades.UseDB().Transaction(func(tx *gorm.DB) error {
		countries := []models.Country{
			{ID: 1, Name: "Afghanistan"},
			{ID: 2, Name: "Albania"},
			{ID: 3, Name: "Algeria"},
			{ID: 4, Name: "American Samoa"},
			{ID: 5, Name: "Andorra"},
			{ID: 6, Name: "Angola"},
			{ID: 7, Name: "Anguilla"},
			{ID: 8, Name: "Antarctica"},
			{ID: 9, Name: "Antigua And Barbuda"},
			{ID: 10, Name: "Argentina"},
			{ID: 11, Name: "Armenia"},
			{ID: 12, Name: "Aruba"},
			{ID: 13, Name: "Australia"},
			{ID: 14, Name: "Austria"},
			{ID: 15, Name: "Azerbaijan"},
			{ID: 16, Name: "Bahamas"},
			{ID: 17, Name: "Bahrain"},
			{ID: 18, Name: "Bangladesh"},
			{ID: 19, Name: "Barbados"},
			{ID: 20, Name: "Belarus"},
			{ID: 21, Name: "Belgium"},
			{ID: 22, Name: "Belize"},
			{ID: 23, Name: "Benin"},
			{ID: 24, Name: "Bermuda"},
			{ID: 25, Name: "Bhutan"},
			{ID: 26, Name: "Bolivia"},
			{ID: 27, Name: "Bosnia And Herzegovina"},
			{ID: 28, Name: "Botswana"},
			{ID: 29, Name: "Bouvet Island"},
			{ID: 30, Name: "Brazil"},
			{ID: 31, Name: "British Indian Ocean Territory"},
			{ID: 32, Name: "Brunei Darussalam"},
			{ID: 33, Name: "Bulgaria"},
			{ID: 34, Name: "Burkina Faso"},
			{ID: 35, Name: "Burundi"},
			{ID: 36, Name: "Cambodia"},
			{ID: 37, Name: "Cameroon"},
			{ID: 38, Name: "Canada"},
			{ID: 39, Name: "Cape Verde"},
			{ID: 40, Name: "Cayman Islands"},
			{ID: 41, Name: "Central African Republic"},
			{ID: 42, Name: "Chad"},
			{ID: 43, Name: "Chile"},
			{ID: 44, Name: "China"},
			{ID: 45, Name: "Christmas Island"},
			{ID: 46, Name: "Cocos (keeling) Islands"},
			{ID: 47, Name: "Colombia"},
			{ID: 48, Name: "Comoros"},
			{ID: 49, Name: "Congo"},
			{ID: 50, Name: "Congo, The Democratic Republic Of The"},
			{ID: 51, Name: "Cook Islands"},
			{ID: 52, Name: "Costa Rica"},
			{ID: 53, Name: "Cote D''ivoire"},
			{ID: 54, Name: "Croatia"},
			{ID: 55, Name: "Cuba"},
			{ID: 56, Name: "Cyprus"},
			{ID: 57, Name: "Czech Republic"},
			{ID: 58, Name: "Denmark"},
			{ID: 59, Name: "Djibouti"},
			{ID: 60, Name: "Dominica"},
			{ID: 61, Name: "Dominican Republic"},
			{ID: 62, Name: "East Timor"},
			{ID: 63, Name: "Ecuador"},
			{ID: 64, Name: "Egypt"},
			{ID: 65, Name: "El Salvador"},
			{ID: 66, Name: "Equatorial Guinea"},
			{ID: 67, Name: "Eritrea"},
			{ID: 68, Name: "Estonia"},
			{ID: 69, Name: "Ethiopia"},
			{ID: 70, Name: "Falkland Islands (malvinas)"},
			{ID: 71, Name: "Faroe Islands"},
			{ID: 72, Name: "Fiji"},
			{ID: 73, Name: "Finland"},
			{ID: 74, Name: "France"},
			{ID: 75, Name: "French Guiana"},
			{ID: 76, Name: "French Polynesia"},
			{ID: 77, Name: "French Southern Territories"},
			{ID: 78, Name: "Gabon"},
			{ID: 79, Name: "Gambia"},
			{ID: 80, Name: "Georgia"},
			{ID: 81, Name: "Germany"},
			{ID: 82, Name: "Ghana"},
			{ID: 83, Name: "Gibraltar"},
			{ID: 84, Name: "Greece"},
			{ID: 85, Name: "Greenland"},
			{ID: 86, Name: "Grenada"},
			{ID: 87, Name: "Guadeloupe"},
			{ID: 88, Name: "Guam"},
			{ID: 89, Name: "Guatemala"},
			{ID: 90, Name: "Guinea"},
			{ID: 91, Name: "Guinea-bissau"},
			{ID: 92, Name: "Guyana"},
			{ID: 93, Name: "Haiti"},
			{ID: 94, Name: "Heard Island And Mcdonald Islands"},
			{ID: 95, Name: "Holy See (vatican City State)"},
			{ID: 96, Name: "Honduras"},
			{ID: 97, Name: "Hong Kong"},
			{ID: 98, Name: "Hungary"},
			{ID: 99, Name: "Iceland"},
			{ID: 100, Name: "India"},
			{ID: 101, Name: "Indonesia"},
			{ID: 102, Name: "Iran, Islamic Republic Of"},
			{ID: 103, Name: "Iraq"},
			{ID: 104, Name: "Ireland"},
			{ID: 105, Name: "Israel"},
			{ID: 106, Name: "Italy"},
			{ID: 107, Name: "Jamaica"},
			{ID: 108, Name: "Japan"},
			{ID: 109, Name: "Jordan"},
			{ID: 110, Name: "Kazakstan"},
			{ID: 111, Name: "Kenya"},
			{ID: 112, Name: "Kiribati"},
			{ID: 113, Name: "Korea, Democratic People''s Republic Of"},
			{ID: 114, Name: "Korea, Republic Of"},
			{ID: 115, Name: "Kosovo"},
			{ID: 116, Name: "Kuwait"},
			{ID: 117, Name: "Kyrgyzstan"},
			{ID: 118, Name: "Lao People''s Democratic Republic"},
			{ID: 119, Name: "Latvia"},
			{ID: 120, Name: "Lebanon"},
			{ID: 121, Name: "Lesotho"},
			{ID: 122, Name: "Liberia"},
			{ID: 123, Name: "Libyan Arab Jamahiriya"},
			{ID: 124, Name: "Liechtenstein"},
			{ID: 125, Name: "Lithuania"},
			{ID: 126, Name: "Luxembourg"},
			{ID: 127, Name: "Macau"},
			{ID: 128, Name: "Macedonia, The Former Yugoslav Republic Of"},
			{ID: 129, Name: "Madagascar"},
			{ID: 130, Name: "Malawi"},
			{ID: 131, Name: "Malaysia"},
			{ID: 132, Name: "Maldives"},
			{ID: 133, Name: "Mali"},
			{ID: 134, Name: "Malta"},
			{ID: 135, Name: "Marshall Islands"},
			{ID: 136, Name: "Martinique"},
			{ID: 137, Name: "Mauritania"},
			{ID: 138, Name: "Mauritius"},
			{ID: 139, Name: "Mayotte"},
			{ID: 140, Name: "Mexico"},
			{ID: 141, Name: "Micronesia, Federated States Of"},
			{ID: 142, Name: "Moldova, Republic Of"},
			{ID: 143, Name: "Monaco"},
			{ID: 144, Name: "Mongolia"},
			{ID: 145, Name: "Montserrat"},
			{ID: 146, Name: "Montenegro"},
			{ID: 147, Name: "Morocco"},
			{ID: 148, Name: "Mozambique"},
			{ID: 149, Name: "Myanmar"},
			{ID: 150, Name: "Namibia"},
			{ID: 151, Name: "Nauru"},
			{ID: 152, Name: "Nepal"},
			{ID: 153, Name: "Netherlands"},
			{ID: 154, Name: "Netherlands Antilles"},
			{ID: 155, Name: "New Caledonia"},
			{ID: 156, Name: "New Zealand"},
			{ID: 157, Name: "Nicaragua"},
			{ID: 158, Name: "Niger"},
			{ID: 159, Name: "Nigeria"},
			{ID: 160, Name: "Niue"},
			{ID: 161, Name: "Norfolk Island"},
			{ID: 162, Name: "Northern Mariana Islands"},
			{ID: 163, Name: "Norway"},
			{ID: 164, Name: "Oman"},
			{ID: 165, Name: "Pakistan"},
			{ID: 166, Name: "Palau"},
			{ID: 167, Name: "Palestinian Territory, Occupied"},
			{ID: 168, Name: "Panama"},
			{ID: 169, Name: "Papua New Guinea"},
			{ID: 170, Name: "Paraguay"},
			{ID: 171, Name: "Peru"},
			{ID: 172, Name: "Philippines"},
			{ID: 173, Name: "Pitcairn"},
			{ID: 174, Name: "Poland"},
			{ID: 175, Name: "Portugal"},
			{ID: 176, Name: "Puerto Rico"},
			{ID: 177, Name: "Qatar"},
			{ID: 178, Name: "Reunion"},
			{ID: 179, Name: "Romania"},
			{ID: 180, Name: "Russian Federation"},
			{ID: 181, Name: "Rwanda"},
			{ID: 182, Name: "Saint Helena"},
			{ID: 183, Name: "Saint Kitts And Nevis"},
			{ID: 184, Name: "Saint Lucia"},
			{ID: 185, Name: "Saint Pierre And Miquelon"},
			{ID: 186, Name: "Saint Vincent And The Grenadines"},
			{ID: 187, Name: "Samoa"},
			{ID: 188, Name: "San Marino"},
			{ID: 189, Name: "Sao Tome And Principe"},
			{ID: 190, Name: "Saudi Arabia"},
			{ID: 191, Name: "Senegal"},
			{ID: 192, Name: "Serbia"},
			{ID: 193, Name: "Seychelles"},
			{ID: 194, Name: "Sierra Leone"},
			{ID: 195, Name: "Singapore"},
			{ID: 196, Name: "Slovakia"},
			{ID: 197, Name: "Slovenia"},
			{ID: 198, Name: "Solomon Islands"},
			{ID: 199, Name: "Somalia"},
			{ID: 200, Name: "South Africa"},
			{ID: 201, Name: "South Georgia And The South Sandwich Islands"},
			{ID: 202, Name: "Spain"},
			{ID: 203, Name: "Sri Lanka"},
			{ID: 204, Name: "Sudan"},
			{ID: 205, Name: "Suriname"},
			{ID: 206, Name: "Svalbard And Jan Mayen"},
			{ID: 207, Name: "Swaziland"},
			{ID: 208, Name: "Sweden"},
			{ID: 209, Name: "Switzerland"},
			{ID: 210, Name: "Syrian Arab Republic"},
			{ID: 211, Name: "Taiwan, Province Of China"},
			{ID: 212, Name: "Tajikistan"},
			{ID: 213, Name: "Tanzania, United Republic Of"},
			{ID: 214, Name: "Thailand"},
			{ID: 215, Name: "Togo"},
			{ID: 216, Name: "Tokelau"},
			{ID: 217, Name: "Tonga"},
			{ID: 218, Name: "Trinidad And Tobago"},
			{ID: 219, Name: "Tunisia"},
			{ID: 220, Name: "Turkey"},
			{ID: 221, Name: "Turkmenistan"},
			{ID: 222, Name: "Turks And Caicos Islands"},
			{ID: 223, Name: "Tuvalu"},
			{ID: 224, Name: "Uganda"},
			{ID: 225, Name: "Ukraine"},
			{ID: 226, Name: "United Arab Emirates"},
			{ID: 227, Name: "United Kingdom"},
			{ID: 228, Name: "United States"},
			{ID: 229, Name: "United States Minor Outlying Islands"},
			{ID: 230, Name: "Uruguay"},
			{ID: 231, Name: "Uzbekistan"},
			{ID: 232, Name: "Vanuatu"},
			{ID: 233, Name: "Venezuela"},
			{ID: 234, Name: "Viet Nam"},
			{ID: 235, Name: "Virgin Islands, British"},
			{ID: 236, Name: "Virgin Islands, U.s."},
			{ID: 237, Name: "Wallis And Futuna"},
			{ID: 238, Name: "Western Sahara"},
			{ID: 239, Name: "Yemen"},
			{ID: 240, Name: "Zambia"},
			{ID: 241, Name: "Zimbabw"},
		}
		for _, country := range countries {
			if err := facades.UseDB().Clauses(clause.OnConflict{DoNothing: true}).Create(&country).Error; err != nil {
				return err
			}
		}

		country := countries[rand.Intn(len(countries))]

		adminHash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		facades.UseDB().Clauses(clause.OnConflict{DoNothing: true}).Create(&models.User{
			AccountName: "Admin",
			Email:       "admin@admin.com",
			Password:    string(adminHash),
			CountryID:   country.ID,
			Country:     country,
		})

		userHash, err := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		users := []*models.User{
			{AccountName: "User", Email: "user@user.com", Password: string(userHash), CountryID: country.ID, Country: country},
			{AccountName: "BR", Email: "br@slc.com", Password: string(userHash), CountryID: country.ID, Country: country},
			{AccountName: "CC", Email: "cc@slc.com", Password: string(userHash), CountryID: country.ID, Country: country},
			{AccountName: "ST", Email: "st@slc.com", Password: string(userHash), CountryID: country.ID, Country: country},
			{AccountName: "VN", Email: "vn@slc.com", Password: string(userHash), CountryID: country.ID, Country: country},
			{AccountName: "TC", Email: "tc@slc.com", Password: string(userHash), CountryID: country.ID, Country: country},
			{AccountName: "LL", Email: "ll@slc.com", Password: string(userHash), CountryID: country.ID, Country: country},
			{AccountName: "GA", Email: "ga@slc.com", Password: string(userHash), CountryID: country.ID, Country: country},
			{AccountName: "JP", Email: "jp@slc.com", Password: string(userHash), CountryID: country.ID, Country: country},
		}

		for _, user := range users {
			facades.UseDB().Clauses(clause.OnConflict{DoNothing: true}).Create(user)
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func runMigration() {
	if err := facades.UseDB().AutoMigrate(
		&models.User{},
		&models.Country{},
		&models.RegisterVerificationToken{},
		&models.Report{},
		&models.UnsuspendRequest{},
	); err != nil {
		log.Fatal(err)
	}
}
