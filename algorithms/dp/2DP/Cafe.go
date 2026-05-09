/*
Около Петиного университета недавно открылось новое кафе, в котором действует следующая система скидок: при каждой покупке более чем на 100 рублей покупатель получает купон, дающий право на один бесплатный обед (при покупке на сумму 100 рублей и меньше такой купон покупатель не получает).

Однажды Пете на глаза попался прейскурант на ближайшие N дней. Внимательно его изучив, он решил, что будет обедать в этом кафе все N дней, причем каждый день он будет покупать в кафе ровно один обед. Однако стипендия у Пети небольшая, и поэтому он хочет по максимуму использовать предоставляемую систему скидок так, чтобы его суммарные затраты были минимальны. Требуется найти минимально возможную суммарную стоимость обедов и номера дней, в которые Пете следует воспользоваться купонами.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1<<20)
  writer := bufio.NewWriterSize(os.Stdout, 1<<20)
  defer writer.Flush()

  line, _ := reader.ReadString('\n')
  n, _ := strconv.Atoi(strings.TrimSpace(line))
  dp, choice := make([][]int, n+1) , make([][]int, n+1)
  for i := 0; i <= n; i++ {
	dp[i] = make([]int, n+1)
	choice[i] = make([]int, n+1)
  }


  for i := 1; i <= n; i++ {dp[0][i] = math.MaxInt}

  for i := 1; i <= n; i++ {
    for j := 0; j <= n; j++ {
        dp[i][j] = math.MaxInt
    }
  }

  nums := make([]int, n)
  for i := 0; i < n; i++{
    line, _ = reader.ReadString('\n')
    curr, _ := strconv.Atoi(strings.TrimSpace(line))
    nums[i] = curr
  }

  for i := 0; i < n; i++{
    cost := nums[i]
    for j := 0; j <= n; j++ {
		  if dp[i][j] == math.MaxInt {
			  continue
		  }
		  if cost <= 100 {
        if dp[i+1][j] > dp[i][j]+cost {
          dp[i+1][j] = dp[i][j] + cost
          choice[i][j] = 0  // купили
        }
      } else {
        if j+1 <= n && dp[i+1][j+1] > dp[i][j]+cost {
          dp[i+1][j+1] = dp[i][j] + cost
          choice[i][j] = 0  // купили
        }
      }
      if j > 0 && dp[i+1][j-1] > dp[i][j] {
        dp[i+1][j-1] = dp[i][j]
       choice[i][j] = 1  // перезаписывает 0 если купон выгоднее!
      }
    }
  }

  	ans, r := math.MaxInt, 0
	for i := 0; i <= n; i++ {
    	if dp[n][i] < ans || (dp[n][i] == ans && i > r) {
        	ans = dp[n][i]
        	r = i
    	}
	}

	usedDays := []int{}
	j := 0
	for i := 0; i < n; i++ {
    	if choice[i][j] == 1 {
        	usedDays = append(usedDays, i+1) // день с 1
        	j--
    	} else {
        	if nums[i] > 100 {
            	j++
        	}
    	}
	}

  writer.WriteString(strconv.Itoa(ans))
  writer.WriteByte('\n')
  writer.WriteString(fmt.Sprintf("%s %s", strconv.Itoa(r), strconv.Itoa(len(usedDays))))
  writer.WriteByte('\n')
  for _, v := range usedDays{
	writer.WriteString(strconv.Itoa(v))
  	writer.WriteByte('\n')
  }
}