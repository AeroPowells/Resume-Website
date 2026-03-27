import { useResumeSection } from '../hooks/useResume'
import styles from './Section.module.css'

export default function Experience() {
  const { data, loading, error } = useResumeSection('/api/resume/experience')

  if (loading) return <p className={styles.state}>Loading...</p>
  if (error) return <p className={styles.state}>Error: {error}</p>

  return (
    <section className={styles.section}>
      <h2 className={styles.heading}>Experience</h2>
      {data.map((job, i) => (
        <div key={i} className={styles.item}>
          <div className={styles.itemHeader}>
            <div>
              <h3 className={styles.itemTitle}>{job.role}</h3>
              <p className={styles.itemSubtitle}>{job.company} &mdash; {job.location}</p>
            </div>
            <span className={styles.dates}>{job.start_date} – {job.end_date}</span>
          </div>
          <ul className={styles.list}>
            {job.highlights.map((h, j) => <li key={j}>{h}</li>)}
          </ul>
        </div>
      ))}
    </section>
  )
}
