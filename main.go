package main

import (
	"fmt"
	"math"
)

// Problem 2.1: List the elements of each of the following sample spaces:
// (a) the set of integers between 1 and 50 divisible by 8;
// (b) the set S = {x | x^2 + 4x - 5=0};
// (c) the set of outcomes when a coin is tossed until a tail or three heads appear;
// (d) the set S = {x | x is a continent};
// (e) the set S = {x | 2x−4 ≥ 0 and x<1}.
func problem2_1() {
	// (a) Set of integers between 1 and 50 divisible by 8
	// We iterate through numbers from 1 to 50 and check if they are divisible by 8.
	// If divisible, we print the number.
	for i := 1; i <= 50; i++ {
		if i%8 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	// (b) Set S = {x | x^2 + 4x - 5 = 0}
	// To find the solutions for x^2 + 4x - 5 = 0, we solve the quadratic equation.
	// The solutions are the values of x that satisfy the equation.
	// In this case, the solutions are -5 and 1.
	fmt.Println("{-5, 1}")

	// (c) Set of outcomes when a coin is tossed until a tail or three heads appear;
	// We list the possible outcomes of tossing a coin until a tail or three heads appear.
	// Here, H represents Head and T represents Tail.
	fmt.Println("{H, TH, TTH, HTTH}")

	// (d) Set S = {x | x is a continent}
	// We list the names of continents.
	fmt.Println("List of continents")

	// (e) Set S = {x | 2x−4 ≥ 0 and x<1}
	// To find the elements of this set, we solve the inequality 2x - 4 ≥ 0.
	// Solving it, we get x ≥ 2. But since x < 1, there's no solution for this set.
	fmt.Println("Empty set")
}

// Problem 2.5: An experiment consists of tossing a die and then flipping a coin once if the number on the die is even.
// If the number on the die is odd, the coin is flipped twice.
// Using the notation 4H to denote the outcome that the die comes up 4 and then the coin comes up heads, and 3HT to denote
// the outcome that the die comes up 3 followed by a head and then a tail on the coin, construct a tree diagram to show
// the 18 elements of the sample space S.
func problem2_5() {
	// Constructing a tree diagram for all possible outcomes.
	fmt.Println("Tree Diagram for Sample Space S:")
	fmt.Println("Die:")
	for _, outcome := range []int{1, 2, 3, 4, 5, 6} {
		fmt.Printf("\t%d\n", outcome)
		fmt.Println("\tCoin:")
		if outcome%2 == 0 {
			fmt.Println("\t\tH")
		} else {
			fmt.Println("\t\tHT")
		}
	}
}

// Problem 2.6: Two jurors are selected from 4 alternates to serve at a murder trial.
// Using the notation A1A3, for example, to denote the simple event that alternates 1 and 3 are selected, list the 6 elements of the sample space S.
func problem2_6() {
	// We list all possible combinations of selecting two jurors from 4 alternates.
	fmt.Println("{A1A2, A1A3, A1A4, A2A3, A2A4, A3A4}")
}

// Problem 2.8: Skipping coding as it involves listing elements corresponding to events and constructing Venn diagram.

// Problem 2.10: An engineering firm is hired to determine if certain waterways in Virginia are safe for fishing.
// Samples are taken from three rivers.
// (a) List the elements of a sample space S, using the letters F for safe to fish and N for not safe to fish.
// (b) List the elements of S corresponding to event E that at least two of the rivers are safe for fishing.
// (c) Define an event that has as its elements the points {FFF, NFF, FFN, NFN}.
func problem2_10() {
	// (a) Sample space S contains all possible combinations of safety statuses for three rivers.
	fmt.Println("S = {FFF, FFN, FNF, FNN, NFF, NFN, NNF, NNN}")

	// (b) Event E contains all outcomes where at least two rivers are safe for fishing.
	fmt.Println("E = {FFF, FFN, FNF, NFF}")

	// (c) This event represents specific combinations of safety statuses for the three rivers.
	fmt.Println("Event with elements: {FFF, NFF, FFN, NFN}")
}

// Problem 2.21: Registrants at a large convention are offered 6 sightseeing tours on each of 3 days.
// In how many ways can a person arrange to go on a sightseeing tour planned by this convention?
func problem2_21() {
	// Total number of sightseeing tours available
	toursPerDay := 6
	days := 3

	// Total number of arrangements
	totalArrangements := int(math.Pow(float64(toursPerDay), float64(days)))
	fmt.Println("Total arrangements:", totalArrangements)
}

// Problem 2.25: A certain brand of shoes comes in 5 different styles, with each style available in 4 distinct colors.
// If the store wishes to display pairs of these shoes showing all of its various styles and colors, how many different pairs will the store have on display?
func problem2_25() {
	styles := 5
	colors := 4

	// Total number of pairs
	totalPairs := styles * colors
	fmt.Println("Total pairs:", totalPairs)
}

// Problem 2.26: A California study concluded that following 7 simple health rules can extend a man’s life by 11 years on average and a woman’s life by 7 years.
// These 7 rules are as follows: no smoking, get regular exercise, use alcohol only in moderation, get 7 to 8 hours of sleep, maintain proper weight, eat breakfast, and do not eat between meals.
// (a) If the person presently violates all 7 rules?
// (b) If the person never drinks and always eats breakfast?
func problem2_26() {
	// (a) Violating all 7 rules
	//allRules := 7
	manLifeExtension := 11
	womanLifeExtension := 7
	totalLifeExtensionAllRules := manLifeExtension + womanLifeExtension
	fmt.Println("Total life extension when violating all rules:", totalLifeExtensionAllRules, "years")

	// (b) Never drinks and always eats breakfast
	// Since the rules that the person follows don't affect the life extension, the total remains the same.
	fmt.Println("Total life extension when never drinking and always eating breakfast:", totalLifeExtensionAllRules, "years")
}

// Problem 2.29: In a fuel economy study, each of 3 race cars is tested using 5 different brands of gasoline at 7 test sites located in different regions of the country.
// If 2 drivers are used in the study, and test runs are made once under each distinct set of conditions, how many test runs are needed?
func problem2_29() {
	raceCars := 3
	gasolineBrands := 5
	testSites := 7
	drivers := 2

	// Total number of test runs needed
	totalTestRuns := raceCars * gasolineBrands * testSites * drivers
	fmt.Println("Total test runs needed:", totalTestRuns)
}

// Problem 2.37: In how many ways can 4 boys and 5 girls sit in a row if the boys and girls must alternate?
func problem2_37() {
	boys := 4
	girls := 5

	// Total number of arrangements
	totalArrangements := factorial(boys) * factorial(girls)
	fmt.Println("Total arrangements with alternating boys and girls:", totalArrangements)
}

// Helper function to calculate factorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Problem 2.49: Find the errors in each of the following statements
func problem2_49() {
	// (a) The probabilities must sum up to 1, but in this case, they sum up to 1.01, indicating an error.
	// (b) The sum of the probabilities of rain and no rain should be 1, but in this case, it sums up to 0.92, indicating an error.
	fmt.Println("Error in statement (a): Probabilities sum up to more than 1")
	fmt.Println("Error in statement (b): Probabilities sum up to less than 1")
}

// Problem 2.51: A box contains 500 envelopes, of which 75 contain $100 in cash, 150 contain $25, and 275 contain $10.
// An envelope may be purchased for $25. What is the sample space for the different amounts of money?
// Assign probabilities to the sample points and then find the probability that the first envelope purchased contains less than $100.
func problem2_51() {
	totalEnvelopes := 500
	//cash100 := 75
	cash25 := 150
	cash10 := 275

	// Sample space: {100, 25, 10}
	// Probabilities: {75/500, 150/500, 275/500}
	//prob100 := float64(cash100) / float64(totalEnvelopes)
	prob25 := float64(cash25) / float64(totalEnvelopes)
	prob10 := float64(cash10) / float64(totalEnvelopes)

	// Probability that the first envelope purchased contains less than $100
	probLessThan100 := prob25 + prob10
	fmt.Println("Probability that the first envelope purchased contains less than $100:", probLessThan100)
}

// Problem 2.52: Suppose that in a senior college class of 500 students...
// If a member of this senior class is selected at random, find the probability that the student
// (a) smokes but does not drink alcoholic beverages;
// (b) eats between meals and drinks alcoholic beverages but does not smoke;
// (c) neither smokes nor eats between meals.
func problem2_52() {
	totalStudents := 500
	smoke := 210
	drink := 258
	eat := 216
	smokeAndDrink := 122
	eatAndDrink := 83
	smokeAndEat := 97
	allThree := 52

	// (a) Smokes but does not drink alcoholic beverages
	probSmokes := float64(smoke) / float64(totalStudents)
	probDrinks := float64(drink) / float64(totalStudents)
	probSmokesNoDrink := probSmokes - float64(smokeAndDrink)/float64(totalStudents)
	fmt.Println("Probability that the student smokes but does not drink alcoholic beverages:", probSmokesNoDrink)

	// (b) Eats between meals and drinks alcoholic beverages but does not smoke
	probEatNoSmokeDrink := (float64(eat) - float64(smokeAndEat)/float64(totalStudents) - float64(eatAndDrink)/float64(totalStudents) + float64(allThree)/float64(totalStudents)) / float64(totalStudents)
	fmt.Println("Probability that the student eats between meals and drinks alcoholic beverages but does not smoke:", probEatNoSmokeDrink)

	// (c) Neither smokes nor eats between meals
	probNeither := 1.0 - probSmokes - probDrinks + (float64(smokeAndDrink) / float64(totalStudents)) + (float64(eatAndDrink) / float64(totalStudents)) - (float64(allThree) / float64(totalStudents))
	fmt.Println("Probability that the student neither smokes nor eats between meals:", probNeither)
}

// Problem 2.53: The probability that an American industry will locate in Shanghai, China, is 0.7...
// What is the probability that the industry will locate
// (a) in both cities?
// (b) in neither city?
func problem2_53() {
	probShanghai := 0.7
	probBeijing := 0.4
	probShanghaiOrBeijing := 0.8

	// (a) In both cities (intersection)
	probBoth := probShanghai + probBeijing - probShanghaiOrBeijing
	fmt.Println("Probability that the industry will locate in both cities:", probBoth)

	// (b) In neither city
	probNeither := 1 - probShanghaiOrBeijing
	fmt.Println("Probability that the industry will locate in neither city:", probNeither)
}

// Problem 2.56: An automobile manufacturer is concerned about a possible recall of its best-selling four-door sedan...
// (a) What is the probability that the defect is the brakes or the fueling system if the probability of defects in both systems simultaneously is 0.15?
// (b) What is the probability that there are no defects in either the brakes or the fueling system?
func problem2_56() {
	probBrakes := 0.25
	probFuel := 0.17
	probBoth := 0.15

	// (a) Probability that the defect is the brakes or the fueling system
	probBrakesOrFuel := probBrakes + probFuel - probBoth
	fmt.Println("Probability that the defect is the brakes or the fueling system:", probBrakesOrFuel)

	// (b) Probability that there are no defects in either the brakes or the fueling system
	probNoDefects := 1 - probBrakesOrFuel
	fmt.Println("Probability that there are no defects in either the brakes or the fueling system:", probNoDefects)
}

// Problem 2.58: A pair of fair dice is tossed. Find the probability of getting
// (a) a total of 8;
// (b) at most a total of 5.
func problem2_58() {
	// Total number of outcomes
	totalOutcomes := 6 * 6

	// (a) Probability of getting a total of 8
	countTotal8 := 5 // Possible combinations: {(2,6), (3,5), (4,4), (5,3), (6,2)}
	probTotal8 := float64(countTotal8) / float64(totalOutcomes)
	fmt.Println("Probability of getting a total of 8:", probTotal8)

	// (b) Probability of getting at most a total of 5
	countTotal5 := 10 // Possible combinations: {(1,1), (1,2), (2,1), (1,3), (2,2), (3,1), (1,4), (2,3), (3,2), (4,1)}
	probTotal5 := float64(countTotal5) / float64(totalOutcomes)
	fmt.Println("Probability of getting at most a total of 5:", probTotal5)
}

// Problem 2.59: In a poker hand consisting of 5 cards, find the probability of holding
// (a) 3 aces;
// (b) 4 hearts and 1 club.
func problem2_59() {
	// Total number of possible poker hands
	totalHands := combinations(52, 5)

	// (a) Probability of holding 3 aces
	count3Aces := combinations(4, 3) * combinations(48, 2)
	prob3Aces := float64(count3Aces) / float64(totalHands)
	fmt.Println("Probability of holding 3 aces:", prob3Aces)

	// (b) Probability of holding 4 hearts and 1 club
	count4Hearts1Club := combinations(13, 4) * combinations(13, 1)
	prob4Hearts1Club := float64(count4Hearts1Club) / float64(totalHands)
	fmt.Println("Probability of holding 4 hearts and 1 club:", prob4Hearts1Club)
}

// Function to calculate combinations (n choose k)
func combinations(n, k int) int {
	if k > n/2 {
		k = n - k
	}
	numerator := 1
	denominator := 1
	for i := 1; i <= k; i++ {
		numerator *= n + 1 - i
		denominator *= i
	}
	return numerator / denominator
}

// Problem 2.73: If R is the event that a convict committed armed robbery and D is the event that the convict pushed dope...
func problem2_73() {
	// (a) P(R|D): Probability of committing armed robbery given that the convict pushed dope
	fmt.Println("(a) P(R|D): Probability of committing armed robbery given that the convict pushed dope")
	// (b) P(D|R): Probability of pushing dope given that the convict committed armed robbery
	fmt.Println("(b) P(D|R): Probability of pushing dope given that the convict committed armed robbery")
	// (c) P(R|D): Probability of committing armed robbery given that the convict pushed dope
	fmt.Println("(c) P(R|D): Probability of committing armed robbery given that the convict pushed dope")
}

// Problem 2.74: A class in advanced physics is composed of 10 juniors, 30 seniors, and 10 graduate students...
func problem2_74() {
	// Given data
	juniors := 10
	seniors := 30
	graduates := 10
	totalStudents := juniors + seniors + graduates
	fmt.Println(totalStudents)
	// Grades data
	juniorsA := 3
	seniorsA := 10
	graduatesA := 5
	totalA := juniorsA + seniorsA + graduatesA

	// Probability that a student chosen at random has earned an A
	//probA := float64(totalA) / float64(totalStudents)

	// Probability that the student is a senior given that he/she has earned an A
	probSeniorGivenA := float64(seniorsA) / float64(totalA)
	fmt.Println("Probability that the student is a senior given that he/she has earned an A:", probSeniorGivenA)
}

// Problem 2.75: A random sample of 200 adults are classified below by sex and their level of education attained...
func problem2_75() {
	// Given data
	//totalMale := 38 + 28 + 22
	//totalFemale := 45 + 50 + 17
	totalSecondary := 28 + 50

	// Probability that the person is a male given that he/she has a secondary education
	probMaleGivenSecondary := float64(28) / float64(totalSecondary)
	fmt.Println("Probability that the person is a male given that he/she has a secondary education:", probMaleGivenSecondary)

	// Probability that a nonsmoker is experiencing no hypertension
	probNonsmokerGivenNoHypertension := float64(38+28+22) / float64(38+28+22+45+50+17)
	fmt.Println("Probability that a nonsmoker is experiencing no hypertension:", probNonsmokerGivenNoHypertension)
}

// Problem 2.76: In an experiment to study the relationship of hypertension and smoking habits...
func problem2_76() {
	// Given data
	totalHeavySmokers := 30 + 19
	totalHeavySmokersHypertension := 26

	// Probability that the person is experiencing hypertension given that he/she is a heavy smoker
	probHypertensionGivenHeavySmoker := float64(totalHeavySmokersHypertension) / float64(totalHeavySmokers)
	fmt.Println("Probability that the person is experiencing hypertension given that he/she is a heavy smoker:", probHypertensionGivenHeavySmoker)
}

// Problem 2.77: In the senior year of a high school graduating class of 100 students...
func problem2_77() {
	// Given data
	totalStudents := 100
	totalMathHistory := 22
	totalMathPsychology := 25
	//totalMathHistoryPsychology := 10

	// (a) Probability that a person enrolled in psychology takes all three subjects
	probEnrolledInAll := float64(totalMathPsychology) / float64(totalStudents)
	fmt.Println("(a) Probability that a person enrolled in psychology takes all three subjects:", probEnrolledInAll)

	// (b) Probability that a person not taking psychology is taking both history and mathematics
	probNotTakingPsychology := float64(totalStudents-totalMathPsychology) / float64(totalStudents)
	probTakingBothMathHistory := float64(totalMathHistory) / float64(totalStudents-totalMathPsychology)
	fmt.Println("(b) Probability that a person not taking psychology is taking both history and mathematics:", probTakingBothMathHistory/probNotTakingPsychology)
}

// Problem 2.78: A manufacturer of a flu vaccine is concerned about the quality of its flu serum...
func problem2_78() {
	// (a) What is the probability that a batch of serum survives the first departmental inspection but is rejected by the second department?
	probFirstSurvivesSecondRejected := 0.9 * 0.08
	fmt.Println("(a) Probability that a batch of serum survives the first departmental inspection but is rejected by the second department:", probFirstSurvivesSecondRejected)

	// (b) What is the probability that a batch of serum is rejected by the third department?
	probThirdRejected := 0.12
	fmt.Println("(b) Probability that a batch of serum is rejected by the third department:", probThirdRejected)
}

// Problem 2.79: In USA Today (Sept. 5, 1996), the results of a survey involving the use of sleepwear while traveling...
func problem2_79() {
	// (a) Probability that a traveler is a female who sleeps in the nude
	probFemaleNude := 0.087
	fmt.Println("(a) Probability that a traveler is a female who sleeps in the nude:", probFemaleNude)

	// (b) Probability that a traveler is male
	probMale := 0.244
	fmt.Println("(b) Probability that a traveler is male:", probMale)

	// (c) Probability that a traveler is male if the traveler sleeps in pajamas or a T-shirt
	probMalePajamasOrTShirt := 0.160 + 0.102
	probMaleGivenPajamasOrTShirt := 0.160 / probMalePajamasOrTShirt
	fmt.Println("(c) Probability that a traveler is male if the traveler sleeps in pajamas or a T-shirt:", probMaleGivenPajamasOrTShirt)
}

// Problem 2.80: The probability that an automobile being filled with gasoline also needs an oil change is 0.25...
func problem2_80() {
	// (a) If the oil has to be changed, what is the probability that a new oil filter is needed?
	// P(OilChange and NewFilter) = P(OilChange) * P(NewFilter|OilChange)
	probOilChange := 0.25
	probNewFilterGivenOilChange := 0.14 / probOilChange
	fmt.Println("(a) Probability that a new oil filter is needed if the oil has to be changed:", probNewFilterGivenOilChange)

	// (b) If a new oil filter is needed, what is the probability that the oil has to be changed?
	// P(OilChange|NewFilter) = P(OilChange and NewFilter) / P(NewFilter)
	probNewFilter := 0.40
	probOilChangeGivenNewFilter := (0.14 / probOilChange) / probNewFilter
	fmt.Println("(b) Probability that the oil has to be changed if a new oil filter is needed:", probOilChangeGivenNewFilter)
}

// Problem 2.81: The probability that a married man watches a certain television show is 0.4...
func problem2_81() {
	// Given data
	probManWatchesShow := 0.4
	probWomanWatchesShow := 0.5
	probManWatchesGivenWomanWatches := 0.7

	// (b) A wife watches the show, given that her husband does
	// P(Watches|Watches) = P(Watches and Watches) / P(Watches)
	probBothWatchShow := probManWatchesGivenWomanWatches * probWomanWatchesShow
	probWifeWatchesGivenHusbandWatches := probBothWatchShow / probManWatchesShow
	fmt.Println("(b) Probability that a wife watches the show, given that her husband does:", probWifeWatchesGivenHusbandWatches)

	// (c) At least one member of a married couple will watch the show
	// P(At least one watches) = 1 - P(Neither watches)
	probNeitherWatches := (1 - probManWatchesShow) * (1 - probWomanWatchesShow)
	probAtLeastOneWatches := 1 - probNeitherWatches
	fmt.Println("(c) Probability that at least one member of a married couple will watch the show:", probAtLeastOneWatches)
}

// Problem 2.82: For married couples living in a certain suburb...
func problem2_82() {
	// Given data
	probHusbandVotes := 0.21
	probWifeVotes := 0.28
	probBothVote := 0.15

	// (a) At least one member of a married couple will vote
	probAtLeastOneVotes := probHusbandVotes + probWifeVotes - probBothVote
	fmt.Println("(a) Probability that at least one member of a married couple will vote:", probAtLeastOneVotes)

	// (b) A wife will vote, given that her husband will vote
	probWifeVotesGivenHusbandVotes := probBothVote / probHusbandVotes
	fmt.Println("(b) Probability that a wife will vote, given that her husband will vote:", probWifeVotesGivenHusbandVotes)

	// (c) A husband will vote, given that his wife will not vote
	probHusbandVotesGivenWifeNotVotes := (probHusbandVotes - probBothVote) / (1 - probWifeVotes)
	fmt.Println("(c) Probability that a husband will vote, given that his wife will not vote:", probHusbandVotesGivenWifeNotVotes)
}

// Problem 2.83: The probability that a vehicle entering the Luray Caverns...
func problem2_83() {
	// Given data
	probCanadianPlates := 0.12
	probCamper := 0.28
	probCanadianPlatesAndCamper := 0.09

	// (a) A camper entering the Luray Caverns has Canadian license plates
	probCamperGivenCanadianPlates := probCanadianPlatesAndCamper / probCanadianPlates
	fmt.Println("(a) Probability that a camper entering the Luray Caverns has Canadian license plates:", probCamperGivenCanadianPlates)

	// (b) A vehicle with Canadian license plates entering the Luray Caverns is a camper
	probCanadianPlatesGivenCamper := probCanadianPlatesAndCamper / probCamper
	fmt.Println("(b) Probability that a vehicle with Canadian license plates entering the Luray Caverns is a camper:", probCanadianPlatesGivenCamper)

	// (c) A vehicle entering the Luray Caverns does not have Canadian plates or is not a camper
	probNotCanadianPlatesOrNotCamper := 1 - (probCanadianPlates + probCamper - probCanadianPlatesAndCamper)
	fmt.Println("(c) Probability that a vehicle entering the Luray Caverns does not have Canadian plates or is not a camper:", probNotCanadianPlatesOrNotCamper)
}

func main() {
	problem2_1()
	problem2_5()
	problem2_6()
	problem2_10()
}
