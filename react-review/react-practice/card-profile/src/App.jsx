function App() {
  return (
    <div className="card">
      <Avatar />
      <div className="data">
        <Intro />
        <SkillList />
      </div>
    </div>
  );
}

function Avatar() {
  return <img className="avatar" src="/artorias.jpg" alt="artorias" />;
}

function Intro() {
  return (
    <div>
      <h1>Artorias, The Abysswalker</h1>
      <p>
        Thou shalt see further on, an Abyss was begat of the ancient beast, and
        threatens to swallow the whole of Oolacile. Knight Artorias came to stop
        this, but such a hero has nary a murmur of Dark. Without doubt he will
        be swallowed by the Abyss, overcome by its utter blackness... Indeed,
        the Abyss may be unstoppable...
      </p>
    </div>
  );
}

function SkillList() {
  return (
    <div className="skill-list">
      <Skill skill="Charging Slash 🗡️" color="orange" />
      <Skill skill=" Abyss Sludge 💥 " color="green" />
      <Skill skill=" Spinning Slash ⚔️" color="red" />
      <Skill skill=" Wrath of the Abyss ⚡" color="blue" />
    </div>
  );
}

function Skill(props) {
  return (
    <div className="skill" style={{ backgroundColor: props.color }}>
      <p>{props.skill}</p>
    </div>
  );
}

export default App;
