// Copyright © 2019 Andreas Fritzler <andreas.fritzler@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var facts = []string{
	"Chuck Norris threw a grenade and killed 50 people, then it exploded.",
	"Chuck Norris can kill two stones with one bird.",
	"Once a cobra bit Chuck Norris' leg. After five days of excruciating pain, the cobra died.",
	"When Chuck Norris was in middle school, his English teacher assigned an essay: 'What is courage?' He received an A+ for turning in a blank page with only his name at the top.",
	"Chuck Norris can pick oranges from an apple tree and make the best lemonade you've ever tasted.",
	"Chuck Norris can kill your imaginary friends.",
	"Chuck Norris can hear sign language.",
	"Chuck Norris once went to mars. Thats why there are no signs of life.",
	"Chuck Norris makes onions cry.",
	"Chuck Norris knows Victoria's secret.",
	"Chuck Norris counted to infinity. Twice.",
	"Chuck Norris beat the sun in a staring contest.",
	"When Bruce Banner gets mad he turns into the Hulk. When the Hulk gets mad he turns into Chuck Norris. When Chuck Norris gets mad, run.",
	"Chuck Norris is the reason Waldo is hiding.",
	"Chuck Norris's Blood Type is AK-47.",
	"Brett Favre can throw a football over 50 yards. Chuck Norris can throw Brett Favre even further.",
	"It is considered a great accomplishment to go down Niagara Falls in a wooden barrel. Chuck Norris can go up Niagara Falls in a cardboard box.",
	"When Chuck Norris enters a room, he doesn't turn the lights on, he turns the dark off.",
	"M.C. Hammer learned the hard way that Chuck Norris can touch this.",
	"Chuck Norris can build a snowman out of rain.",
	"Chuck Norris can strangle you with a cordless phone.",
	"Chuck Norris was once charged with three attempted murders in Boulder County, but the Judge quickly dropped the charges because Chuck Norris does not 'attempt' murder.",
	"Bill Gates lives in constant fear that Chuck Norris' PC will crash.",
	"Death once had a near-Chuck-Norris experience.",
	"Chuck Norris was once on Celebrity Wheel of Fortune and was the first to spin. The next 29 minutes of the show consisted of everyone standing around awkwardly, waiting for the wheel to stop.",
	"Chuck Norris' dog is trained to pick up his own poop because Chuck Norris will not take shit from anyone.",
	"When Chuck Norris gives you the finger, he's telling you how many seconds you have left to live.",
	"Giraffes were created when Chuck Norris uppercutted a horse.",
	"If it looks like chicken, tastes like chicken, and feels like chicken but Chuck Norris says its beef, then it's beef.",
	"When the Boogeyman goes to sleep every night he checks his closet for Chuck Norris.",
	"Chuck Norris doesn't cheat death. He wins fair and square.",
	"Leading hand sanitizers claim they can kill 99.9 percent of germs. Chuck Norris can kill 100 percent of whatever the hell he wants.",
	"Chuck Norris puts the 'laughter' in 'manslaughter'.",
	"Chuck actually died four years ago, but the Grim Reaper can't get up the courage to tell him.",
	"There is no theory of evolution, just a list of creatures Chuck Norris allows to live.",
	"Chuck Norris' calendar goes straight from March 31st to April 2nd. No one fools Chuck Norris.",
	"Chuck Norris doesn't have good aim. His bullets just know better than to miss.",
	"When a zombie apocalypse starts, Chuck Norris doesn't try to survive. The zombies do.",
	"Chuck Norris will never have a heart attack... even a heart isnt foolish enough to attack Chuck Norris.",
	"Chuck Norris beat Halo 1, 2, and 3 on Legendary with a broken Guitar Hero controller.",
	"Chuck can set ants on fire with a magnifying glass. At night.",
	"Chuck Norris's daughter lost her virginity, he got it back.",
	"The real reason Hitler killed himself is because he found out that Chuck Norris is Jewish.",
	"Chuck Norris tells Simon what to do.",
	"The reason the Holy Grail has never been recovered is because nobody is brave enough to ask Chuck Norris to give up his favourite coffee mug.",
	"Chuck Norris sleeps with a pillow under his gun.",
	"Chuck Norris doesn't play 'hide-and-seek.' He plays 'hide-and-pray-I-don't-find-you.'",
	"Chuck Norris is the only person that can punch a cyclops between the eye.",
	"Chuck Norris plays russian roulette with a fully loded revolver... and wins.",
	"Chuck Norris is the only person on the planet that can kick you in the back of the face.",
	"Chuck Norris can speak braille.",
	"Chuck Norris can do a wheelie on a unicycle.",
	"Chuck refers to himself in the fourth person.",
	"Chuck Norris is the only man to ever defeat a brick wall in a game of tennis.",
	"Chuck Norris doesnt wear a watch. He decides what time it is.",
	"Chuck Norris once bowled a perfect game with a marble.",
	"Chuck Norris's computer has no 'backspace' button, Chuck Norris doesn't make mistakes.",
	"Chuck Norris can hit you so hard your blood will bleed.",
	"Jack was nimble, Jack was quick, but Jack still couldn't dodge Chuck Norris' roundhouse kick.",
	"Chuck Norris can drown a fish.",
	"Chuck Norris can unscramble an egg.",
	"Chuck Norris' tears cure cancer. Too bad he has never cried.",
	"When Chuck Norris was a freshman in High School, 7 seniors took him Snipe hunting in the woods. He killed 19 Snipes, a grizzly bear and 3 timber wolves. Oddly, the 7 seniors are still missing.",
	"Chuck Norris has a diary. It's called the Guinness Book of World Records.",
	"Chuck Norris can literally throw a party.",
	"Chuck Norris CAN find the end of a circle.",
	"The saddest moment for a child is not when he learns Santa Claus isn't real, it's when he learns Chuck Norris is.",
	"The only time Chuck Norris was wrong was when he thought he had made a mistake.",
	"If you spell Chuck Norris wrong on Google it doesn't say, 'Did you mean Chuck Norris?' It simply replies, 'Run while you still have the chance.'",
	"Some kids piss their name in the snow. Chuck Norris can piss his name into concrete.",
	"Chuck Norris doesn't dial the wrong number, you pick up the wrong phone.",
	"Why hasn't a video game been made about Chuck Norris? Simple: nobody controls Chuck Norris.",
	"Chuck Norris can cut a knife with butter.",
	"Chuck Norris can speak French... In Russian.",
	"Chuck Norris can make a slinky go upstairs.",
	"If he wanted to, Chuck Norris could rob a bank. By phone.",
	"A bulletproof vest wears Chuck Norris for protection.",
	"The original title for Alien vs. Predator was Alien and Predator vs Chuck Norris. The film was cancelled shortly after going into preproduction. No one would pay nine dollars to see a movie fourteen seconds long.",
	"The reason newborn babies cry is because they know they have just entered a world with Chuck Norris.",
	"On a high school math test, Chuck Norris put down 'Violence' as every one of the answers. He got an A+ on the test because Chuck Norris solves all his problems with Violence.",
	"Chuck can divide by zero.",
	"A Handicapped parking sign does not signify that this spot is for handicapped people. It is actually in fact a warning, that the spot belongs to Chuck Norris and that you will be handicapped if you park there.",
	"Police label anyone attacking Chuck Norris as a Code 45-11.... a suicide.",
	"Chuck Norris narrates Morgan Freeman's life.",
	"Every Chuck Norris joke is a five star joke just because it says Chuck Norris.",
	"Chuck Norris can sit in the corner of a round room.",
	"If you swallow a quarter and Chuck Norris round house kicks you in the stomach you will crap out two dimes and a nickel.",
	"Tough men eat nails. Chuck Norris does all his grocery shopping at Home Depot.",
	"Chuck Norris can run a 3-legged race by himself.",
	"Chuck Norris doesn't wear a condom because theres no such thing as protection from Chuck Norris.",
	"Big foot claims he saw Chuck Norris.",
	"Chuck Norris doesn't worry about high gas prices. His vehicles run on fear.",
	"Chuck Norris doesn't teabag , he potato sacks.",
	"Chuck Norris used to beat the shit out of his shadow because it was following to close. It now stands a safe 30 feet behind him.",
	"Chuck Norris sleeps with a night light. Not because Chuck Norris is afraid of the dark, but the dark is afraid of Chuck Norris.",
	"Chuck Norris does not hunt because the word hunting implies the possibility of failure. Chuck Norris goes killing.",
	"Chuck Norris can delete the Recycling Bin.",
	"Chuck Norris CAN believe it's not butter.",
	"The Great Wall of China was originally created to keep Chuck Norris out. It failed miserably.",
	"Chuck Norris can start a fire by staring at wood.",
	"Chuck Norris doesn't need containers, he ships with a round-house kick.",
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a Chuck Norris fact",
	Long:  `Get a Chuck Norris fact.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		idx := r.Intn(len(facts))
		fmt.Println(facts[idx])
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
