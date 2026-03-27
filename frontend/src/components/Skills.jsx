import { useResumeSection } from '../hooks/useResume'
import styles from './Section.module.css'
import skillStyles from './Skills.module.css'

export default function Skills() {
  const { data, loading, error } = useResumeSection('/api/resume/skills')

  if (loading) return <p className={styles.state}>Loading...</p>
  if (error) return <p className={styles.state}>Error: {error}</p>

  return (
    <section className={styles.section}>
      <h2 className={styles.heading}>Skills</h2>
      <div className={skillStyles.grid}>
        {data.map((group, i) => (
          <div key={i} className={skillStyles.group}>
            <h3 className={skillStyles.category}>{group.category}</h3>
            <div className={skillStyles.tags}>
              {group.skills.map((skill) => (
                <span key={skill} className={skillStyles.tag}>{skill}</span>
              ))}
            </div>
          </div>
        ))}
      </div>
    </section>
  )
}
