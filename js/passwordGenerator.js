const LOWERCASE = 'abcdefghijklmnopqrstuvwxyz'
const UPPERCASE = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
const DIGITS = '0123456789'
const SPECIAL = '!@#$%^&*'

const nextRandom = (seed) => (16807 * seed) % 2147483647

const generatePassword = (length, seed, options = {}) => {
  const { useUppercase = true, useDigits = true, useSpecial = false } = options

  let alphabet = LOWERCASE
  if (useUppercase) alphabet += UPPERCASE
  if (useDigits) alphabet += DIGITS
  if (useSpecial) alphabet += SPECIAL

  // Math.abs перед %, т.к. в JS % отрицательного числа = отрицательный результат
  let current = Math.abs(seed) % 2147483647
  if (current === 0) current = 1

  let result = ''
  for (let i = 0; i < length; i += 1) {
    current = nextRandom(current)
    result += alphabet[current % alphabet.length]
  }
  return result
}

const hasLowercase = (password) => {
  for (const char of password) {
    if (char >= 'a' && char <= 'z') return true
  }
  return false
}

const hasUppercase = (password) => {
  for (const char of password) {
    if (char >= 'A' && char <= 'Z') return true
  }
  return false
}

const hasDigit = (password) => {
  for (const char of password) {
    if (char >= '0' && char <= '9') return true
  }
  return false
}

const hasSpecial = (password) => {
  for (const char of password) {
    if (SPECIAL.includes(char)) return true
  }
  return false
}

const isLongEnough = (password, minLength = 8) => password.length >= minLength

const strengthScore = (password) => {
  let score = 0
  if (isLongEnough(password)) score += 1
  if (hasLowercase(password)) score += 1
  if (hasUppercase(password)) score += 1
  if (hasDigit(password)) score += 1
  if (hasSpecial(password)) score += 1
  return score
}

const checkPassword = (password) => {
  const score = strengthScore(password)
  let verdict
  if (score <= 2) {
    verdict = 'Слабый'
  } else if (score === 3) {
    verdict = 'Средний'
  } else if (score === 4) {
    verdict = 'Надёжный'
  } else {
    verdict = 'Очень надёжный'
  }
  return `${verdict} пароль (оценка ${score} из 5)`
}

export { generatePassword, checkPassword }
