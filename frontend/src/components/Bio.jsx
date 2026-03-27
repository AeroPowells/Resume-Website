import { useResumeSection } from '../hooks/useResume'
import styles from './Bio.module.css'

export default function Bio() {
  const { data, loading, error } = useResumeSection('/api/resume/bio')

  if (loading) return <p className={styles.state}>Loading...</p>
  if (error) return <p className={styles.state}>Error: {error}</p>

  return (
    <header className={styles.bio}>
      <h1 className={styles.name}>{data.name}</h1>
      <p className={styles.title}>{data.title}</p>
      <p className={styles.summary}>{data.summary}</p>
      <div className={styles.contact}>
        <span>{data.email}</span>
        <span>{data.phone}</span>
        <span>{data.location}</span>
      </div>
      <div className={styles.links}>
        {data.links?.map((link) => (
          <a key={link.label} href={link.url} target="_blank" rel="noreferrer">
            {link.label}
          </a>
        ))}
      </div>
    </header>
  )
}
