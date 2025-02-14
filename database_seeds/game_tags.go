package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm/clause"
)

func SeedGameTags() error {
	for _, x := range []models.GameTag{
		{Name: "Action"},
		{Name: "Low Confidence Metric"},
		{Name: "Adventure"},
		{Name: "Casual"},
		{Name: "Singleplayer"},
		{Name: "Simulation"},
		{Name: "Strategy"},
		{Name: "RPG"},
		{Name: "2D"},
		{Name: "Atmospheric"},
		{Name: "Free to Play"},
		{Name: "Puzzle"},
		{Name: "Early Access"},
		{Name: "Multiplayer"},
		{Name: "Story Rich"},
		{Name: "Fantasy"},
		{Name: "Violent"},
		{Name: "Anime"},
		{Name: "Great Soundtrack"},
		{Name: "Pixel Graphics"},
		{Name: "Nudity"},
		{Name: "VR"},
		{Name: "Sexual Content"},
		{Name: "First-Person"},
		{Name: "Funny"},
		{Name: "Gore"},
		{Name: "Sci-fi"},
		{Name: "Arcade"},
		{Name: "Cute"},
		{Name: "Horror"},
		{Name: "Shooter"},
		{Name: "Colorful"},
		{Name: "Difficult"},
		{Name: "Platformer"},
		{Name: "Sports"},
		{Name: "Retro"},
		{Name: "Exploration"},
		{Name: "3D"},
		{Name: "Family Friendly"},
		{Name: "Open World"},
		{Name: "Relaxing"},
		{Name: "Co-op"},
		{Name: "Survival"},
		{Name: "Female Protagonist"},
		{Name: "Massively Multiplayer"},
		{Name: "Racing"},
		{Name: "Visual Novel"},
		{Name: "FPS"},
		{Name: "Comedy"},
		{Name: "Turn-Based"},
		{Name: "Third Person"},
		{Name: "Online Co-Op"},
		{Name: "Action-Adventure"},
		{Name: "Sandbox"},
		{Name: "Realistic"},
		{Name: "Physics"},
		{Name: "Point & Click"},
		{Name: "Space"},
		{Name: "Top-Down"},
		{Name: "Mystery"},
		{Name: "Choices Matter"},
		{Name: "Stylized"},
		{Name: "PvP"},
		{Name: "Psychological Horror"},
		{Name: "Management"},
		{Name: "Adult Content"},
		{Name: "Replay Value"},
		{Name: "Minimalist"},
		{Name: "Building"},
		{Name: "Local Multiplayer"},
		{Name: "Dark"},
		{Name: "Tactical"},
		{Name: "Cartoony"},
		{Name: "Design & Illustration"},
		{Name: "Side Scroller"},
		{Name: "Puzzle-Platformer"},
		{Name: "Multiple Endings"},
		{Name: "Shoot 'Em Up"},
		{Name: "Software"},
		{Name: "Classic"},
		{Name: "Old School"},
		{Name: "2D Platformer"},
		{Name: "Education"},
		{Name: "Party-Based RPG"},
		{Name: "Rogue-like"},
		{Name: "Turn-Based Strategy"},
		{Name: "Character Customization"},
		{Name: "Utilities"},
		{Name: "Survival Horror"},
		{Name: "Local Co-Op"},
		{Name: "Action RPG"},
		{Name: "Mature"},
		{Name: "Movie"},
		{Name: "Fast-Paced"},
		{Name: "Short"},
		{Name: "Rogue-lite"},
		{Name: "Zombies"},
		{Name: "Hand-drawn"},
		{Name: "Controller"},
		{Name: "Procedural Generation"},
		{Name: "Crafting"},
		{Name: "JRPG"},
		{Name: "Hidden Object"},
		{Name: "Combat"},
		{Name: "RPGMaker"},
		{Name: "Historical"},
		{Name: "War"},
		{Name: "Turn-Based Combat"},
		{Name: "Bullet Hell"},
		{Name: "Magic"},
		{Name: "Hack and Slash"},
		{Name: "Memes"},
		{Name: "Experimental"},
		{Name: "Medieval"},
		{Name: "Futuristic"},
		{Name: "Web Publishing"},
		{Name: "Romance"},
		{Name: "Resource Management"},
		{Name: "RTS"},
		{Name: "Cartoon"},
		{Name: "Stealth"},
		{Name: "Fighting"},
		{Name: "PvE"},
		{Name: "3D Platformer"},
		{Name: "Walking Simulator"},
		{Name: "Turn-Based Tactics"},
		{Name: "Music"},
		{Name: "Logic"},
		{Name: "Dungeon Crawler"},
		{Name: "Linear"},
		{Name: "Choose Your Own Adventure"},
		{Name: "Post-apocalyptic"},
		{Name: "Dating Sim"},
		{Name: "Interactive Fiction"},
		{Name: "Top-Down Shooter"},
		{Name: "Dark Fantasy"},
		{Name: "Isometric"},
		{Name: "Drama"},
		{Name: "Beautiful"},
		{Name: "Base Building"},
		{Name: "Surreal"},
		{Name: "Card Game"},
		{Name: "Competitive"},
		{Name: "Military"},
		{Name: "4 Player Local"},
		{Name: "Text-Based"},
		{Name: "Driving"},
		{Name: "Cyberpunk"},
		{Name: "Soundtrack"},
		{Name: "Tower Defense"},
		{Name: "Third-Person Shooter"},
		{Name: "Score Attack"},
		{Name: "Flight"},
		{Name: "Economy"},
		{Name: "Board Game"},
		{Name: "Robots"},
		{Name: "Narration"},
		{Name: "Abstract"},
		{Name: "1990's"},
		{Name: "2.5D"},
		{Name: "Action Roguelike"},
		{Name: "Hentai"},
		{Name: "LGBTQ+"},
		{Name: "Dark Humor"},
		{Name: "City Builder"},
		{Name: "Aliens"},
		{Name: "1980s"},
		{Name: "Metroidvania"},
		{Name: "Character Action Game"},
		{Name: "Detective"},
		{Name: "Perma Death"},
		{Name: "Moddable"},
		{Name: "Animation & Modeling"},
		{Name: "Beat 'em up"},
		{Name: "Team-Based"},
		{Name: "Tabletop"},
		{Name: "Investigation"},
		{Name: "Arena Shooter"},
		{Name: "Time Management"},
		{Name: "NSFW"},
		{Name: "Thriller"},
		{Name: "Nature"},
		{Name: "Cinematic"},
		{Name: "Twin Stick Shooter"},
		{Name: "Level Editor"},
		{Name: "Real Time Tactics"},
		{Name: "World War II"},
		{Name: "Clicker"},
		{Name: "Emotional"},
		{Name: "Automobile Sim"},
		{Name: "Psychological"},
		{Name: "Audio Production"},
		{Name: "Strategy RPG"},
		{Name: "Wargame"},
		{Name: "Immersive Sim"},
		{Name: "Destruction"},
		{Name: "Game Development"},
		{Name: "Tutorial"},
		{Name: "Psychedelic"},
		{Name: "Precision Platformer"},
		{Name: "Grand Strategy"},
		{Name: "Tactical RPG"},
		{Name: "Crime"},
		{Name: "Life Sim"},
		{Name: "Addictive"},
		{Name: "Demons"},
		{Name: "Rhythm"},
		{Name: "Conversation"},
		{Name: "Match 3"},
		{Name: "Dystopian"},
		{Name: "MMORPG"},
		{Name: "2D Fighter"},
		{Name: "Loot"},
		{Name: "Video Production"},
		{Name: "Alternate History"},
		{Name: "Parkour"},
		{Name: "Comic Book"},
		{Name: "Modern"},
		{Name: "Runner"},
		{Name: "Episodic"},
		{Name: "Open World Survival Craft"},
		{Name: "Mouse only"},
		{Name: "Nonlinear"},
		{Name: "Space Sim"},
		{Name: "Software Training"},
		{Name: "Real-Time"},
		{Name: "Dark Comedy"},
		{Name: "Science"},
		{Name: "CRPG"},
		{Name: "Lovecraftian"},
		{Name: "Supernatural"},
		{Name: "Inventory Management"},
		{Name: "Souls-like"},
		{Name: "Split Screen"},
		{Name: "Blood"},
		{Name: "Steampunk"},
		{Name: "Lore-Rich"},
		{Name: "Cult Classic"},
		{Name: "Artificial Intelligence"},
		{Name: "Voxel"},
		{Name: "Mythology"},
		{Name: "Noir"},
		{Name: "Philosophical"},
		{Name: "Mechs"},
		{Name: "Battle Royale"},
		{Name: "Swordplay"},
		{Name: "Grid-Based Movement"},
		{Name: "Deckbuilding"},
		{Name: "eSports"},
		{Name: "Dragons"},
		{Name: "Cats"},
		{Name: "4X"},
		{Name: "Card Battler"},
		{Name: "Trains"},
		{Name: "Tanks"},
		{Name: "Political"},
		{Name: "Pirates"},
		{Name: "Photo Editing"},
		{Name: "6DOF"},
		{Name: "Parody"},
		{Name: "Hex Grid"},
		{Name: "Real-Time with Pause"},
		{Name: "Otome"},
		{Name: "Remake"},
		{Name: "Bullet Time"},
		{Name: "Colony Sim"},
		{Name: "Co-op Campaign"},
		{Name: "Hacking"},
		{Name: "Collectathon"},
		{Name: "Trading"},
		{Name: "Ninja"},
		{Name: "Satire"},
		{Name: "Naval"},
		{Name: "Western"},
		{Name: "Farming Sim"},
		{Name: "Time Manipulation"},
		{Name: "Word Game"},
		{Name: "Class-Based"},
		{Name: "MOBA"},
		{Name: "Time Travel"},
		{Name: "Mining"},
		{Name: "Vehicular Combat"},
		{Name: "GameMaker"},
		{Name: "Agriculture"},
		{Name: "Dinosaurs"},
		{Name: "America"},
		{Name: "Hunting"},
		{Name: "FMV"},
		{Name: "3D Vision"},
		{Name: "Gothic"},
		{Name: "Dungeons & Dragons"},
		{Name: "Politics"},
		{Name: "Touch-Friendly"},
		{Name: "God Game"},
		{Name: "Kickstarter"},
		{Name: "Gun Customization"},
		{Name: "Capitalism"},
		{Name: "Illuminati"},
		{Name: "Programming"},
		{Name: "Underwater"},
		{Name: "3D Fighter"},
		{Name: "Mystery Dungeon"},
		{Name: "Idler"},
		{Name: "Superhero"},
		{Name: "Martial Arts"},
		{Name: "Documentary"},
		{Name: "Trading Card Game"},
		{Name: "Combat Racing"},
		{Name: "Experience"},
		{Name: "Conspiracy"},
		{Name: "Fishing"},
		{Name: "Automation"},
		{Name: "Solitaire"},
		{Name: "Cold War"},
		{Name: "Hero Shooter"},
		{Name: "Sokoban"},
		{Name: "Spectacle fighter"},
		{Name: "Minigames"},
		{Name: "Dog"},
		{Name: "Vampire"},
		{Name: "Underground"},
		{Name: "Political Sim"},
		{Name: "Assassin"},
		{Name: "Time Attack"},
		{Name: "Heist"},
		{Name: "Faith"},
		{Name: "Dynamic Narration"},
		{Name: "Soccer"},
		{Name: "Naval Combat"},
		{Name: "Football"},
		{Name: "Epic"},
		{Name: "Quick-Time PrivateChatSockets"},
		{Name: "Mod"},
		{Name: "Asynchronous Multiplayer"},
		{Name: "TrackIR"},
		{Name: "Diplomacy"},
		{Name: "Looter Shooter"},
		{Name: "Music-Based Procedural Generation"},
		{Name: "Typing"},
		{Name: "Chess"},
		{Name: "Party Game"},
		{Name: "Archery"},
		{Name: "On-Rails Shooter"},
		{Name: "Sailing"},
		{Name: "Party"},
		{Name: "Villain Protagonist"},
		{Name: "Transportation"},
		{Name: "Narrative"},
		{Name: "Crowdfunded"},
		{Name: "Immersive"},
		{Name: "360 Video"},
		{Name: "Trivia"},
		{Name: "Gaming"},
		{Name: "Pinball"},
		{Name: "Sequel"},
		{Name: "Offroad"},
		{Name: "Based On A Novel"},
		{Name: "Silent Protagonist"},
		{Name: "Creature Collector"},
		{Name: "Horses"},
		{Name: "Mars"},
		{Name: "World War I"},
		{Name: "Sniper"},
		{Name: "Games Workshop"},
		{Name: "Snow"},
		{Name: "Foreign"},
		{Name: "LEGO"},
		{Name: "Rome"},
		{Name: "Auto Battler"},
		{Name: "Boxing"},
		{Name: "Warhammer 40K"},
		{Name: "Gambling"},
		{Name: "Traditional Roguelike"},
		{Name: "Roguelike Deckbuilder"},
		{Name: "Unforgiving"},
		{Name: "Transhumanism"},
		{Name: "Medical Sim"},
		{Name: "Golf"},
		{Name: "Motorbike"},
		{Name: "Werewolves"},
		{Name: "Electronic Music"},
		{Name: "Farming"},
		{Name: "Nostalgia"},
		{Name: "Bikes"},
		{Name: "Asymmetric VR"},
		{Name: "Basketball"},
		{Name: "Spelling"},
		{Name: "Mini Golf"},
		{Name: "Ambient"},
		{Name: "Roguevania"},
		{Name: "Cooking"},
		{Name: "Intentionally Awkward Controls"},
		{Name: "Jet"},
		{Name: "Outbreak Sim"},
		{Name: "Submarine"},
		{Name: "Social Deduction"},
		{Name: "Pool"},
		{Name: "Spaceships"},
		{Name: "Wrestling"},
		{Name: "Lemmings"},
		{Name: "Feature Film"},
		{Name: "Baseball"},
		{Name: "Tennis"},
		{Name: "Instrumental Music"},
		{Name: "Benchmark"},
		{Name: "Skateboarding"},
		{Name: "Hockey"},
		{Name: "Voice Control"},
		{Name: "Motocross"},
		{Name: "Electronic"},
		{Name: "Skating"},
		{Name: "Cycling"},
		{Name: "Skiing"},
		{Name: "Rock Music"},
		{Name: "Bowling"},
		{Name: "Hardware"},
		{Name: "Steam Machine"},
		{Name: "Reboot"},
		{Name: "Snowboarding"},
		{Name: "8-bit Music"},
		{Name: "Well-Written"},
		{Name: "BMX"},
		{Name: "ATV"},
		{Name: "Action RTS"},
	} {
		if err := facades.UseDB().Clauses(clause.OnConflict{DoNothing: true}).Create(&x).Error; err != nil {
			return err
		}
	}

	return nil
}
