package game

type Command int8

// a : Add Line - Adds one broken line to the bottom of the target field.
// c : Clear Line - Clears the line on the bottom of the target field.
// b : Clear Special Blocks - Causes all special blocks on the target field to return to normal blocks.
// r : Random Blocks Clear - Random blocks are removed from the target field, often creating gaps in lines.
// o : Block Bomb - Causes any Block Bombs on the target field to "explode", causing surrounding pieces to scatter throughout the field.
// q : Blockquake - Causes blocks on each line on the target field to shift, creating an earthquake-like effect.
// g : Block Gravity - Causes blocks on that target field to immediately fall into any gaps and instantly deletes any complete lines which are created.
// s : Switch Fields - User and target switch fields.
// n : Nuke Field - Removes all blocks from the target field.
// Immunity - Target cannot be affected by any special for a default of 15 seconds.
// Clear Column - Deletes a random column from the target field.
// Mutate Pieces - A certain number of the target's next pieces (default 3) become large and awkwardly shaped.
// Darkness - Target cannot see anything on his field besides the current piece and its immediate surroundings for a default of 10 seconds.
// Confusion - Causes target's controls to "mix up" for a default of 10 seconds (for example, the key that normally moves the piece left may start moving the piece right instead.)
