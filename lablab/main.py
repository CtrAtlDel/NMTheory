import numpy as np

lines = 0
column = 0


def read_matrix_from_file(filename):
    with open(filename, 'r') as file:
        rows = int(file.readline().strip())
        columns = int(file.readline().strip())

        lines = rows
        columns = columns

        matrix = np.empty((rows, columns), dtype=int)
        for i in range(rows):
            line = file.readline().strip()
            row_values = np.fromstring(line, sep=',', dtype=int)
            matrix[i] = row_values

    return matrix, lines, columns


def create_identity_matrix(size):
    matrix = np.eye(size)
    return matrix


def concatenate_zero_row(matrix):
    zero_row = np.zeros((1, matrix.shape[1]), dtype=matrix.dtype)
    concatenated_matrix = np.vstack((matrix, zero_row))
    transposed_matrix = np.transpose(concatenated_matrix)
    return transposed_matrix


def concatenate_matrix(matrix1, matrix2):
    return np.vstack((matrix1, matrix2))


def get_matrix_a(matrix):  # Формирование матрицы А
    lines, cols = matrix.shape
    matrix[:, -1] = -matrix[:, -1]
    half_matrix = concatenate_zero_row(create_identity_matrix(cols - 1))
    matrix_a = concatenate_matrix(matrix, half_matrix)
    return matrix_a


def get_matrix_a_shtrix(matrix_a):
    matrix = matrix_a
    line, column = matrix_a.shape
    for j in range(0, column - 1):
        while check_zeroes(j, j, matrix):
            matrix = change_matrix_upper_left_corner(matrix, j, j)  # Перемещаем в левый угол наименьший элемент шаг 1
            for i in range(j + 1, column):  # Потому что исходный столбец не трогаем, шаг 2
                d = matrix[j, i] // matrix[j, j]
                matrix = subtract_columns(matrix, i, j, d)
    return matrix


def check_zeroes(line_index, col_index, matrix):
    line_elements = matrix[line_index, col_index + 1:]
    return np.all(line_elements != 0)


def divide_row_by_value(matrix, row_index, d):
    matrix[row_index] //= d
    return matrix


def subtract_columns(matrix, col_index_from, col_index_subtract, multiplier):
    matrix[:, col_index_from] -= multiplier * matrix[:, col_index_subtract]
    return matrix


def change_matrix_upper_left_corner(matrix, row_index, element_index):
    nonzero_indices = np.nonzero(matrix[row_index])[0]
    smallest_nonzero = np.abs(matrix[row_index, nonzero_indices]).min()
    smallest_index = np.abs(matrix[row_index, element_index:]).argmin() + element_index

    matrix[:, [element_index, smallest_index]] = matrix[:, [smallest_index, element_index]]
    if smallest_nonzero < 0:
        matrix[:, element_index] = -matrix[:, element_index]

    return matrix


def remove_last_column(matrix):
    return matrix[:, :-1]


def get_last_column(matrix):
    return matrix[:, -1]


def concatenate_with_last_column(matrix, column):
    column = column.reshape(-1, 1)

    concatenated = np.concatenate((matrix, column), axis=1)

    return concatenated


def get_matrix_result(matrix, line_n):
    l = []
    matrx_result = matrix
    line, col = matrix.shape
    for i in range(0, line_n):
        k = matrix[i, -1] // matrix[i, i]
        matrx_result = subtract_columns(matrx_result, -1, i, k)
    return matrx_result


def check(matrix, line):
    for i in range(0, line):
        if matrix[i, -1] != 0:
            print("Cannot be solve")
            break


def print_column_from_element(matrix, start_line_index, column_index):  # Печать К матрицы
    column = matrix[start_line_index:, column_index]
    for element in column:
        print(element)


filename = './data.txt'

if __name__ == '__main__':
    matrix, line, column = read_matrix_from_file(filename)

    matrix_a = get_matrix_a(matrix)
    line_a, col_a = matrix_a.shape
    last_coll = get_last_column(matrix_a)
    matrx_a_shtrix = get_matrix_a_shtrix(remove_last_column(matrix_a))
    new_matrix_with_b = concatenate_with_last_column(matrx_a_shtrix, last_coll)
    matrix_result = get_matrix_result(new_matrix_with_b, line)
    print(matrx_a_shtrix)
    print(new_matrix_with_b)
    print(matrix_result)
    check(matrix_result, line)
    for i in range(0, col_a):
        print_column_from_element(matrix_result, i, i)
