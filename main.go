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

func check_win(board[][]string,piece string) int{
	// UP DOWN POSITION

	win := false
	for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
			if i + 3 < 6 {
				if board[5-i][j] == piece && board[4-i][j] == piece && board[3-i][j] == piece && board[2-i][j] == piece {
					win = true
				}
			}
        }
    }

	// RIGHT LEFT POSITION

	for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
			if j + 3 < 7 {
				if board[i][6-j] == piece && board[i][5-j] == piece && board[i][4-j] == piece && board[i][3-j] == piece {
					win = true
				}
			}
        }
    }
	
	// DIAG RIGHT POSITION

	for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
			if j + 3 < 7 && i + 3 < 6{
				if board[5-i][j] == piece && board[4-i][j+1] == piece && board[3-i][j+2] == piece && board[2-i][j+3] == piece {
					win = true
				}
			}
        }
    }

	// DIAG LEFT POSITION

	for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
			if j + 3 < 7 && i + 3 < 7{
				if board[5-i][6-j] == piece && board[4-i][5-j] == piece && board[3-i][4-j] == piece && board[2-i][3-j] == piece {
					win = true
				}
			}
        }
    }

	if win == true && piece == "X"{
		return 1
	} else if win == true && piece == "O"{
		return -1
	} else {
		return 0
	}
}

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
		for x < 0 || y < 0 {
			x,y = right_row_ia(board)
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

	check_all(board)
	return(board)
}


func check_all(board[][]string) {
	check_w := 0
	check_l := 0
	
	check_w = check_win(board,"X")
	check_l = check_win(board,"O")
		
	if check_w == 1 {
		fmt.Println("YOU WIN STEEVEN")
		os.Exit(0)
	} else if check_l == -1 {
		fmt.Println("YOU WIN IA")
		os.Exit(0)
	}
	if check_full(board) == true {
		fmt.Println("FULLL")
		os.Exit(0)
	}
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
