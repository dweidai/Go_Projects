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
	for i := 0; i < numEntries; i++ {
		listens := rand.Intn(1000)
		likes := 0
		if(listens != 0){
			likes = rand.Intn(listens)
		}
		comments := 0
		if(likes != 0){
			comments = rand.Intn(likes)
		}
		randomTime := rand.Int63n(time.Now().Unix() - 94608000) + 94608000
		uploadTime := time.Unix(randomTime, 0)
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

// Attempted using Wilson score interval but twisted with golden ratio and like/views ratio
// high confidence score should come to the front of the viewers
func confidence(likes int, listens int) float64{
	if (likes == 0){
		return 0
	}
	like := float64(likes)
	listen := float64(listens)
	z := 1.61803398875
	p := like /listen
	left := p+1.0/(2.0*(listen))*z*z
	right := z*math.Sqrt(p*(1.0-p)/(listen) + z*z/(4.0*listen*listen))
	under := 1.0+1.0/(listen)*z*z

	return (left-right)/under
}

// combination of Clopper-Pearson interval with sigmoid 
// take as a penalty, if the the penalty score is low then the confidence is high
func discriminant(likes int, comments int, diffs int) float64{
	active := float64(likes) + float64(comments)
	if (diffs == 0){
		diffs = 1
	}
	diff := float64(diffs)
	p := (active/diff)*(1-(active/diff))/diff
	//diff = math.Log(float64(diffs))
	z := 1.61803398875
	upper := (1- (p + (z*z)/2/diff))
	lower := 1 + (z*z)/diff
	x := 1- 1 / (1+math.Exp(upper/lower * (-1)))
	return x-0.2
}
/** 
 * Calculate the time difference between the uploaded time and current time
 */
func timeDiff(uploadTime time.Time) int{
	duration := time.Since(uploadTime)
	minDiff := duration.Minutes()
	return int(minDiff)
}

/**
 * 
 * 
 */
func score(audio Audio) float64{
	conf := confidence(audio.likes, audio.listens)
	diff := timeDiff(audio.uploadTime)
	disc := discriminant(audio.likes, audio.comments, diff)
	if conf-disc <= 0{
		return 0
	}
	return conf-disc
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
	db := randomDataBase(50)
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