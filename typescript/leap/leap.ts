const isLeapYear = (year: number): boolean => (year % 100 == 0 ? year % 400 == 0 : year % 4 == 0)

export default isLeapYear
