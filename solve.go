package main

import "fmt"
import "strconv"

type numSeries []int

type cell struct {
    candidatesRow   numSeries
    candidatesCol   numSeries
    candidatesField numSeries
    value           int
    candidateValue  int
}

func (cell *cell) nextCandidate() int {
    offset := 0
    if cell.candidateValue > 0 {
        offset = cell.candidateValue - 1
        cell.candidatesRow[offset] = cell.candidateValue
        cell.candidatesCol[offset] = cell.candidateValue
        cell.candidatesField[offset] = cell.candidateValue
        cell.candidateValue = 0
        offset += 1
    }
    for i := offset; i < 9; i++ {
        if cell.candidatesRow[i] > 0 &&
            cell.candidatesCol[i] > 0 &&
            cell.candidatesField[i] > 0 {
                cell.candidateValue = cell.candidatesRow[i]
                cell.candidatesRow[i] = 0
                cell.candidatesCol[i] = 0
                cell.candidatesField[i] = 0
                return cell.candidateValue
        }
    }
    return 0
}

func (cell *cell) getValue() int {
    if cell.value > 0 {
        return cell.value
    }

    if cell.candidateValue > 0 {
        return cell.candidateValue
    }

    return 0
}

type solvingBoard [81]cell

func (board *solvingBoard) print() {
    for i, cell := range board {
        row := i / 9
        col := i % 9
        if col == 0 && (row == 0 || row == 3 || row == 6) {
            fmt.Println("-------------------------")
        }
        if col == 0 || col == 3 || col == 6 {
            fmt.Print("| ")
        }
        fmt.Print(cell.getValue(), " ")
        if col == 8 {
            fmt.Println("|")
        }
    }
    fmt.Println("-------------------------")
}

func (board *solvingBoard) solve(pos int) bool {
    if pos == len(board) {
        return board[len(board) - 1].getValue() > 0
    }

    cell := &board[pos]

    if cell.getValue() > 0 {
        return board.solve(pos + 1)
    }

    for cell.nextCandidate() > 0 {
        if board.solve(pos + 1) == true {
            return true
        }
    }

    return false
}

func solveSudoku(board [][]byte) {
    sudoku := createSudoku(board)

    sudoku.print()

    solved := sudoku.solve(0)

    fmt.Println("Solved:", solved)

    sudoku.print()
}

func createSudoku(board [][]byte) solvingBoard {
    cols := make([]numSeries, 9)
    rows := make([]numSeries, 9)
    fields := make([]numSeries, 9)

    for i := 0; i < 9; i++ {
        cols[i] = make(numSeries, 9)
        rows[i] = make(numSeries, 9)
        fields[i] = make(numSeries, 9)

        for j := 0; j < 9; j++ {
            cols[i][j] = j + 1
            rows[i][j] = j + 1
            fields[i][j] = j + 1
        }
    }

    cells := solvingBoard{}

    for i := 0; i < 81; i++ {
        row := i / 9
        col := i % 9
        field := row / 3 * 3 + col / 3
        cells[i] = cell{
            candidatesRow: rows[row],
            candidatesCol: cols[col],
            candidatesField: fields[field],
        }
        if board[row][col] == '.' {
            cells[i].value = 0
        } else {
            v, _ := strconv.Atoi(string(board[row][col]))
            cells[i].value = v
            cells[i].candidatesRow[v-1] = 0
            cells[i].candidatesCol[v-1] = 0
            cells[i].candidatesField[v-1] = 0
        }
    }

    return cells
}

func main() {
    fmt.Println("Solving sudoku...")
    board := [][]byte{
        {'.','.','9','7','4','8','.','.','.'},
        {'7','.','.','.','.','.','.','.','.'},
        {'.','2','.','1','.','9','.','.','.'},
        {'.','.','7','.','.','.','2','4','.'},
        {'.','6','4','.','1','.','5','9','.'},
        {'.','9','8','.','.','.','3','.','.'},
        {'.','.','.','8','.','3','.','2','.'},
        {'.','.','.','.','.','.','.','.','6'},
        {'.','.','.','2','7','5','9','.','.'},
    }
    // board := [][]byte{
    //     {'5','3','.','.','7','.','.','.','.'},
    //     {'6','.','.','1','9','5','.','.','.'},
    //     {'.','9','8','.','.','.','.','6','.'},
    //     {'8','.','.','.','6','.','.','.','3'},
    //     {'4','.','.','8','.','3','.','.','1'},
    //     {'7','.','.','.','2','.','.','.','6'},
    //     {'.','6','.','.','.','.','2','8','.'},
    //     {'.','.','.','4','1','9','.','.','5'},
    //     {'.','.','.','.','8','.','.','7','9'},
    // }
    solveSudoku(board)
}
