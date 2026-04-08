Banner Loader Utility

A lightweight Go function that transforms a standard banner text file into a searchable map. It maps ASCII characters (32 to 126) to their corresponding graphical string slices.
⚙️ Features

    ASCII Mapping: Automatically maps decimal characters 32 (Space) through 126 (~).

    Structured Output: Returns a map[rune][]string for O(1) lookup time.

    Efficient Parsing: Uses os.ReadFile and strings.Split for fast data ingestion.

🛠 Usage

To use this utility, ensure you have a banner file where each character occupies exactly 9 lines (1 header/empty line + 8 lines of character art).

📂 Data Structure

The function expects a specific file format where characters are vertically stacked.

Component	Description

Filename	Path to the .txt file containing banner art.

Rune (Key)	The ASCII character (e.g., 'H').

[]string (Value)	The 8 lines of text that form the visual character.

⚠️ Error Handling

    File Issues: If the file is missing or unreadable, the function prints file could not be opened to the console and returns nil.

    Guard Clause: The function utilizes the "return early" pattern to prevent unnecessary processing upon failure.

📝 Best Practices Used

    Idiomatic Returns: Returns (nil) on failure to signal the caller that the map is invalid.

    Casting: Properly handles rune and int conversions for ASCII math.

    Slicing: Uses efficient memory slicing lines[start+1 : start+9] instead of copying data.