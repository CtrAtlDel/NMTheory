import numpy as np

def change_matrix_upper_left_corner(matrix, row_index, element_index):
    # Get the indices of the non-zero elements in the specified row
    nonzero_indices = np.nonzero(matrix[row_index])[0]

    # Find the smallest absolute non-zero element in the specified row
    smallest_nonzero = np.abs(matrix[row_index, nonzero_indices]).min()

    # Find the index of the smallest non-zero element
    smallest_index = np.abs(matrix[row_index, element_index:]).argmin() + element_index

    # Swap the columns to move the smallest non-zero element to the leftmost column
    matrix[:, [element_index, smallest_index]] = matrix[:, [smallest_index, element_index]]

    # Change the signs of elements in the column if the smallest non-zero element is negative
    if smallest_nonzero < 0:
        matrix[:, element_index] = -matrix[:, element_index]

    return matrix

# Example usage
matrix = np.array([[5, -3, 5], [2, 4, -6], [1, 7, 8]])
transformed_matrix = change_matrix_upper_left_corner(matrix, 0, 0)  # Transform the element at row 1, column 1
print(transformed_matrix)

