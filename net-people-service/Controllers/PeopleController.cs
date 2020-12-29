using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using People.Service.Models;

namespace People.Service.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class PeopleController : ControllerBase
    {
        IPeopleProvider provider;

        public PeopleController(IPeopleProvider provider)
        {
            this.provider = provider;
        }

        [HttpGet]
        public async Task<IReadOnlyCollection<Person>> Get()
        {
            await Task.Delay(3000);
            return provider.GetPeople();
        }

        [HttpGet("{id}")]
        public async Task<Person> Get(int id)
        {
            await Task.Delay(1000);
            return provider.GetPeople().FirstOrDefault(p => p.Id == id);
        }

        [HttpGet("ids")]
        public IReadOnlyCollection<int> GetIds()
        {
            return provider.GetPeople().Select(p => p.Id).ToList();
        }
    }
}
