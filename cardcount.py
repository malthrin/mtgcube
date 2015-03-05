		
def printFileContents(fileName, fileContents, verbose):
	count = 0
	lastSubheader = ''
	subheaderCount = 0
	subheaderCounts = [ ]
	for line in fileContents.splitlines():
		if verbose and line.startswith('#'):
			if len(lastSubheader) > 0:
				subheaderCounts.append((lastSubheader, subheaderCount))
			lastSubheader = line.strip('# ')
			subheaderCount = 0
		elif len(line) == 0 or line.startswith('#'):
			continue
		else:
			count += 1
			subheaderCount += 1
			
	print fileName, count
	if verbose:
		if len(lastSubheader) > 0:
			subheaderCounts.append((lastSubheader, subheaderCount))
		for sub, count in subheaderCounts:
			print '\t', sub, count
	
files = [ 'white', 'blue', 'black', 'red', 'green', 'multicolor', 'artifact', 'land' ]

for filename in files:
	with open(filename, 'r') as file:
		printFileContents(filename, file.read(), True)
