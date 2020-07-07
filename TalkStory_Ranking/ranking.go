package main

import (
	"os"
	"fmt"
	"time"
	"math"
	"bufio"
	"math/rand"
)

type Audio struct{
	listens int // number of listens
	likes int // number of likes
	comments int // number of comments 
	uploadTime time.Time // the time this audio is published
	viewed bool // whether the user has viewed it
	liked bool // whether the user has liked it
}

func Print(audio Audio){
	fmt.Println("-------------------")
	fmt.Println("|")
	fmt.Println("| number of listens: ", audio.listens)
	fmt.Println("| number of likes: ", audio.likes)
	fmt.Println("| number of comments: ", audio.comments)
	fmt.Println("| Uploaded time: ", audio.uploadTime)
	if audio.viewed{
		fmt.Println("| Viewed: YES")
	} else{
		fmt.Println("| Viewed: NO")
	}
	if audio.liked{
		fmt.Println("| Liked: YES")
	} else{
		fmt.Println("| Liked: NO")
	}
	fmt.Println("|")
	fmt.Println("-------------------")
}

func newAudio(numListens int, numLikes int, numComments int,  upload time.Time, view bool, like bool) *Audio {
	toReturn := new(Audio)
	toReturn.listens = numListens
	toReturn.likes = numLikes
	toReturn.comments = numComments
	toReturn.uploadTime = upload
	toReturn.viewed = view
	toReturn.liked = like
	return toReturn
}

/*
 * Generate fake data for testing 
 */
func randomDataBase(numEntries int) []Audio{
	toReturn := make([]Audio, numEntries)
	fresh := new(Audio)
	fresh.listens = 0
	fresh.likes = 0
	fresh.comments = 0
	fresh.uploadTime = time. Now()
	fresh.viewed = false
	fresh.liked = false
	toReturn[0] = *fresh
	Print(*fresh)
	for i := 1; i < numEntries; i++ {
		listens := rand.Intn(100)
		likes := 0
		if(listens != 0){
			likes = rand.Intn(listens)
		}
		comments := 0
		if(likes != 0){
			comments = rand.Intn(likes)
		}
		min := time.Date(2020, 6, 20, 0, 0, 0, 0, time.UTC).Unix()
		max := time.Date(2020, 7, 4, 0, 0, 0, 0, time.UTC).Unix()
		delta := max - min
		sec := rand.Int63n(delta) + min
		uploadTime := time.Unix(sec, 0)
		viewed := true 
		if rand.Intn(2) == 1{
			viewed = false
		}
		liked := false
		if rand.Intn(2) == 1 && viewed{
			liked = true
		}
		temp := newAudio(listens, likes, comments, uploadTime, viewed, liked)
		Print(*temp)
		toReturn[i] = *temp
	}
	return toReturn
}

// Attempted using Wilson score interval but twisted with E and like/time ratio
// high confidence score should come to the front of the viewers
func confidence(likes int, comments int, diffs int) float64{
	if likes == 0 {
		return 0
	}
	diff := (float64(diffs))/100
	if (diffs <= 3){
		diffs = 1
	}
	like := float64(likes) + float64(comments) * math.SqrtE
	z := math.E
	p := like /diff
	left := p+z/(2.0*(diff))*z*z
	right := z*math.Sqrt(p*(1.0-p)/(diff) + z*z/(4.0*diff*diff))
	under := 1.0+1.0/(diff)*z*z
	return (left-right)/under
}

// combination of Clopper-Pearson interval with sigmoid 
// take as a penalty, if the the penalty score is low then the confidence is high
func discriminant(likes int, comments int, listens int) float64{
	active := float64(likes) + float64(comments)
	if (listens == 0){
		listens = 1
	}
	listen := float64(listens)
	p := (active/listen)*(1-(active/listen))/listen
	z := math.E
	upper := (1- (p + (z*z)/2/listen))
	lower := 1 + (z*z)/listen
	//sigmoid
	//x := 1- 1 / (1+math.Exp(upper/lower * (-1)))
	return upper/lower
}
/** 
 * Calculate the time difference between the uploaded time and current time
 */
func timeDiff(uploadTime time.Time) int{
	duration := time.Since(uploadTime)
	minDiff := duration.Hours()
	return int(minDiff)
}


func score(audio Audio) float64{
	diff := timeDiff(audio.uploadTime)
	conf := confidence(audio.likes, audio.comments, diff)
	disc := discriminant(audio.likes, audio.comments, audio.listens)
	//if conf-disc <= 0{
	//	return 0
	//}
	return 1+conf-disc
}

/**
 * Merge sort the audios based on the score
 */
func Sort(items []Audio) []Audio{
	var num = len(items)
      
    if num == 1 {
        return items
    }
      
    middle := int(num / 2)
    var (
        left = make([]Audio, middle)
        right = make([]Audio, num-middle)
    )
    for i := 0; i < num; i++ {
        if i < middle {
            left[i] = items[i]
        } else {
            right[i-middle] = items[i]
        }
    }
      
    return merge(Sort(left), Sort(right))
}

/** merge sort part 2 (helper funciton) */
func merge(left, right []Audio) (result []Audio) {
    result = make([]Audio, len(left) + len(right))
      
    i := 0
    for len(left) > 0 && len(right) > 0 {
        if score(left[0]) > score(right[0]) {
            result[i] = left[0]
            left = left[1:]
        } else if(left[0].likes > right[0].likes){
            result[i] = right[0]
            right = right[1:]
        } else{
        	result[i] = right[0]
            right = right[1:]
        }
        i++
    }
      
    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }
      
    return
}
/**
 * Split into three categories
 * 1. Never viewed, never liked
 * 2. Viewed, but did not hit like
 * 3. Viewed, and liked
 */
func split(audio []Audio) (neverViewed []Audio, viewedNoLiked []Audio, viewedLiked []Audio){
	var never []Audio
	var viewed []Audio
	var liked []Audio
	for i := 0; i < len(audio); i++ {
		if (audio[i].viewed == false){
			never = append(never, audio[i])
		} else if (audio[i].liked == false){
			viewed = append(viewed, audio[i])
		} else{
			liked = append(liked, audio[i])
		}
	}
	return never, viewed, liked
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\t HERE are the data")
	db := randomDataBase(30)
	/*fmt.Println("Done generating data, press enter to show sorted")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	if text == "\n" {
		sortDB := Sort(db)
		fmt.Println()
		fmt.Println("Sorted and Ranked")
		fmt.Println()
		for i := 0; i < len(sortDB); i++ {
			Print(sortDB[i])
			fmt.Println(score(sortDB[i]))
		}
	}*/
	never, viewed, liked := split(db)
	fmt.Println("Done generating data, press enter to show sorted category 1")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	if text == "\n"{
		neverSort := Sort(never)
		fmt.Println()
		fmt.Println("Sorted and Ranked")
		fmt.Println()
		for i := 0; i < len(neverSort); i++ {
			Print(neverSort[i])
		}
	} else {
		os.Exit(1)
	}
	
	fmt.Println("Done showing category 1 data, press enter to show sorted category 2")
	text1, _ := reader.ReadString('\n')
	if text1 == "\n"{
		viewedSort := Sort(viewed)
		fmt.Println()
		fmt.Println("Sorted and Ranked")
		fmt.Println()
		for i := 0; i < len(viewedSort); i++ {
			Print(viewedSort[i])
		}
	} else {
		os.Exit(1)
	}
	fmt.Println("Done showing category 2 data, press enter to show sorted category 3")
	text2, _ := reader.ReadString('\n')
	if text2 == "\n"{
		likedSort := Sort(liked)
		fmt.Println()
		fmt.Println("Sorted and Ranked")
		fmt.Println()
		for i := 0; i < len(likedSort); i++ {
			Print(likedSort[i])
		}
	} else {
		os.Exit(1)
	}
}