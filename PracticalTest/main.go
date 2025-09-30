package main

import (
	"PracticalTest/model"
	"encoding/json"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/summary", func(w http.ResponseWriter, r *http.Request) {
		summaries, err := SummaryBooking()
		if err != nil {
			http.Error(w, "Failed to fetch summary", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		result := map[string]interface{}{
			"status":  "success",
			"message": "Summary data retrieved successfully",
			"data":    summaries,
			"Meta": map[string]interface{}{
				"count": len(summaries),
			},
		}
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8080", nil)
}

func SummaryBooking() ([]model.Summary, error) {

	// fetch booking data
    bookings, err := fetchBookingData()
    if err != nil {
        return nil, err
    }

	// fetch type consumption data
    typeConsumptions, err := fetchTypeConsumptionData()
    if err != nil {
        return nil, err
    }

    // create map consumption type for easy lookup
    typeConsumMap := make(map[string]model.TypeConsum)
    totalJenisHarga := 0.0
    for _, tc := range typeConsumptions {
        typeConsumMap[tc.Name] = tc
        totalJenisHarga += tc.MaxPrice
    }

    summaryMap := make(map[string]*model.Summary)
	// process each booking
    for _, booking := range bookings {
        key := booking.OfficeName + "|" + booking.RoomName

        if _, exists := summaryMap[key]; !exists {
            summaryMap[key] = &model.Summary{
                OfficeName:     booking.OfficeName,
                RoomName:       booking.RoomName,
                Capacity:       0,
                UsagePercent:   0,
                ConsumeNominal: 0,
                TypeConsume:    []model.TypeDetail{},
            }
        }

        // add capacity (number of participants)
        summaryMap[key].Capacity += booking.Participants

        // calculate consumption per booking
        consumeCountMap := make(map[string]int)
        for _, consume := range booking.ListConsumption {
            consumeCountMap[consume.Name]++
        }

		// calculate total consumption nominal and type detail
        for consumeType, count := range consumeCountMap {
            if tc, ok := typeConsumMap[consumeType]; ok {
                amount := float64(count) * tc.MaxPrice * float64(booking.Participants)
                summaryMap[key].ConsumeNominal += amount

                found := false
                for i := range summaryMap[key].TypeConsume {
                    if summaryMap[key].TypeConsume[i].Name == consumeType {
                        summaryMap[key].TypeConsume[i].Count += count * booking.Participants
                        summaryMap[key].TypeConsume[i].Amount += amount
                        found = true
                        break
                    }
                }
                if !found {
                    summaryMap[key].TypeConsume = append(summaryMap[key].TypeConsume, model.TypeDetail{
                        Name:   consumeType,
                        Count:  count * booking.Participants,
                        Amount: amount,
                    })
                }
            }
        }
    }

    
   // calculate usage percentage per room
	for _, summary := range summaryMap {
		if summary.Capacity > 0 && totalJenisHarga > 0 {
			maxPossible := float64(summary.Capacity) * totalJenisHarga
			percent := (summary.ConsumeNominal / maxPossible) * 100

			// round to 2 decimal places with math.Round
			summary.UsagePercent = math.Round(percent*100) / 100
			
		}
	}


    // change map to slice for result
    var summaries []model.Summary
    for _, s := range summaryMap {
        summaries = append(summaries, *s)
    }

    return summaries, nil
}




func fetchBookingData() ([]model.Booking, error) {

	bookingList := []model.Booking{}

	// fetch booking data from mock API
	resp, err := http.Get("https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList")
	if err != nil {
		return nil, err
	}
	// defer closing the response body
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&bookingList); err != nil {
		return nil, err
	}
	return bookingList, nil
}

func fetchTypeConsumptionData() ([]model.TypeConsum, error) {
	typeConsumList := []model.TypeConsum{}
	// fetch type consumption data from mock API
	resp, err := http.Get("https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi")
	if err != nil {
		return nil, err
	}

	// defer closing the response body
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&typeConsumList); err != nil {
		return nil, err
	}
	return typeConsumList, nil
}
