package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "math/rand"
//import "sort"
import "time"

const rows = 6
const cols = 7

func make_board() [][]string{
	board := make([][]string, rows)
    for i := range board {
        board[i] = make([]string, cols)
    }
	for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            board[i][j] = "."
        }
    }
	return(board)
}

func print_board(board[][]string) {
	for i := 0; i < rows; i++ {
		for j:= 0; j < cols; j++ {
			fmt.Print(board[i][j])
		}
		fmt.Println()
	}
}


func my_atoi(s string)int {
	x , err := strconv.Atoi(s)
	if err != nil {
		return(-1)
	}
	return (x)
}

func right_row(board[][]string,position int) (int,int) {
	k := 6
	for i := 0; i < rows; i++ {
		for j := 0; j < cols ;j++ {
			if k-j-1 < 0 || position > 7 || position < 1 {
				return -1,-1
			} 
			if board[k-j-1][position - 1] == "." {
				return k-j-1,position-1
			}
		}
	}
	return -1,-1
}

//func monte_carlo(board[]string) int {}

func right_row_ia(board[][]string) (int,int) {
	rand.Seed(time.Now().UnixNano())
	position := rand.Intn(7) + 1
	k := 6
	for i := 0; i < rows; i++ {
		for j := 0; j < cols ;j++ {
			if k-j-1 < 0 || position > 7 || position < 0 {
				return -1,-1
			}
			if board[k-j-1][position - 1] == "." {
				return k-j-1,position-1
			}
		}
	}
	return -1,-1
}

/*fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaa = " , position)

func check_left_right(board[]string,piece string) int{
	win := 0
	all_win := []int{}
	for j := 0; j < 7; j++ {
		idx := 41 - j
		for i := 0; i < 6 ; i++ {
			if board[idx] == piece {
				win = win + 1
				idx = idx - 7
			} else if board[idx] != piece {
				idx = idx - 7
			}
			if i == 5 {		
				all_win = append(all_win,win)
				win = 0
			}
		}
	}
	sort.Ints(all_win)
	sort.Sort(sort.Reverse(sort.IntSlice(all_win)))
	fmt.Println(all_win)
	return(all_win[0])
}

func check_up_down(board[]string,piece string) int{
	win := 0
	//up := 0
	all_win := []int{}
	for j := 0; j < 6; j++ {
		idx := 0 + (j * 6)
		//fmt.Println("idx = " , idx)
		for i := 0; i < 7 ; i++ {
			//fmt.Println("IDX = " , idx)
			if board[idx] == piece {
				win = win + 1
				idx = idx + 7
			} else if board[idx] != piece {
				idx = idx + 7
			}
			if i == 5 {		
				all_win = append(all_win,win)
				win = 0
			}
		}
	}
	sort.Ints(all_win)
	sort.Sort(sort.Reverse(sort.IntSlice(all_win)))
	fmt.Println(all_win)
	return(all_win[0])
}

/*
func check_win(board[]string,int piece) {
	left_right := 0
	up_down := 0
	left_right_diag_up := 0
	left_right_diag_down := 0

	//for i := 0; i < 
	
}
*/

func auto_play(board[][]string,position int,piece string) [][]string {

	x:=0
	y:=0
	for i := 0 ; i < 8 ; i ++ {
		x,y = right_row(board,position+i)
		if x > -1 && y > -1 {
			board[x][y] = piece
			return (board)
		}
	}
	return (board)
}

func check_full(board[][]string) bool {
	for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
        	if board[i][j] == "." {
				board[i][j] = "."
				return (false)
			}
        }
    }
	return (true)
}

func drop_piece(board[][]string,position int,piece string,ia bool) [][]string{
	x,y := 0,0
	
	if ia == true {		
		x,y = right_row_ia(board)
		//fmt.Println(x,y)
		for x < 0 || y < 0 {
			x,y = right_row_ia(board)
			fmt.Println(x,y," pos problem")
			//return(board)
		}
		board[x][y] = piece
		return(board)
	}
	x,y = right_row(board,position)
	if x < 0 || y < 0 {
		board = auto_play(board,position,"X")
		return(board)
	}
	if position == 1 {
		x,y = right_row(board,position)
		board[x][y] = piece
	} else if position == 2 {
		x,y = right_row(board,position)
		board[x][y] = piece
	} else if position == 3 {
		x,y = right_row(board,position)
		board[x][y] = piece
	} else if position == 4 {
		x,y = right_row(board,position)
		board[x][y] = piece
	} else if position == 5 {
		x,y = right_row(board,position)
		board[x][y] = piece
	} else if position == 6 {
		x,y = right_row(board,position)
		board[x][y] = piece
	} else if position == 7 {
		x,y = right_row(board,position)
		board[x][y] = piece
	}
	//fmt.Println("x = " , x , " y = " , y)
	return(board)
}


func main_game() {
	board := make_board()
	position := 0
	ia_play_1 := 0
	ia_play_2 := 0
	play := bufio.NewScanner(os.Stdin)
	r1 := 0
	r2 := 0
	print_board(board)
	fmt.Println("Choose a pos 1-7")
	for play.Scan() {
		position = my_atoi(play.Text())
		ia_play_1 , ia_play_2 = right_row_ia(board)
		r1,r2 = right_row(board,position)
		if check_full(board) == true {
			fmt.Println("FULLL")
			os.Exit(0)
		}
		//fmt.Println(position,ia_play_1,ia_play_2,r1,r2)
		//fmt.Println(ia_play_1,ia_play_2,r1,r2)
		if position != -1 || r1 != -1 || r2 != -1 || ia_play_1 != -1 || ia_play_2 != -1{
			drop_piece(board,position,"X",false)
			drop_piece(board,ia_play_1,"O",true)
			print_board(board)
		}
		fmt.Println("Choose a pos 1-7")

	}
}

func main() {
	main_game()
	//fmt.Println(board)
}