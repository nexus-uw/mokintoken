async function main() {
  if (location.pathname === '/') {
    const { Welcome } = await import('./welcome')
    new Welcome()
  }

  if (location.pathname.startsWith('/decrypt')) {
    const { Decrypt } = await import('./decrypt')
    new Decrypt()
  }
}
main()
