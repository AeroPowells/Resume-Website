import Bio from './components/Bio'
import Experience from './components/Experience'
import Education from './components/Education'
import Skills from './components/Skills'
import styles from './App.module.css'

export default function App() {
  return (
    <main className={styles.layout}>
      <Bio />
      <Experience />
      <Education />
      <Skills />
    </main>
  )
}
